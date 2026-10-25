package channellab

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 超时只负责报告挂死；正确性由 channel 握手、关闭和计数验证。
const guard = 3 * time.Second

func next[T any](t *testing.T, ch <-chan T) (T, bool) {
	t.Helper()
	timer := time.NewTimer(guard)
	defer timer.Stop()
	select {
	case v, ok := <-ch:
		return v, ok
	case <-timer.C:
		t.Fatal("等待 channel 超时：检查是否漏了 close 或取消分支")
		var zero T
		return zero, false
	}
}

func collect[T any](t *testing.T, ch <-chan T) []T {
	t.Helper()
	var values []T
	for {
		v, ok := next(t, ch)
		if !ok {
			return values
		}
		values = append(values, v)
	}
}

func source(values ...int) <-chan int {
	ch := make(chan int, len(values))
	for _, n := range values {
		ch <- n
	}
	close(ch)
	return ch
}

func assertInts(t *testing.T, got, want []int) {
	t.Helper()
	// 这里允许 nil 切片与长度为零的切片等价。
	if len(got) != len(want) {
		t.Fatalf("got %v; want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v; want %v", got, want)
		}
	}
}

func TestE01Sum(t *testing.T) {
	if got := Sum(source(0, 1, -2, 8)); got != 7 {
		t.Fatalf("got %d; want 7", got)
	}
	if got := Sum(source()); got != 0 {
		t.Fatalf("empty: got %d; want 0", got)
	}
}

func TestE02TrySend(t *testing.T) {
	ch := make(chan int, 1)
	if !TrySend(ch, 7) || TrySend(ch, 8) {
		t.Fatal("空缓冲应成功，满缓冲应失败")
	}
	if v, _ := next(t, ch); v != 7 {
		t.Fatalf("got %d; want 7", v)
	}
	if TrySend(nil, 1) {
		t.Fatal("nil channel 的非阻塞发送应失败")
	}
	if TrySend(make(chan int), 1) {
		t.Fatal("没有接收者的无缓冲 channel 应返回 false")
	}
}

func TestE03Receive(t *testing.T) {
	ctx := context.Background()
	in := source(0)
	v, ok, err := Receive(ctx, in)
	if v != 0 || !ok || err != nil {
		t.Fatalf("真实零值: %d %v %v", v, ok, err)
	}
	v, ok, err = Receive(ctx, in)
	if v != 0 || ok || err != nil {
		t.Fatalf("关闭且排空: %d %v %v", v, ok, err)
	}
	for _, mode := range []string{"already-canceled", "blocked-on-nil"} {
		t.Run(mode, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if mode == "already-canceled" {
				cancel()
			}
			result := make(chan error, 1)
			go func() {
				_, _, err := Receive(ctx, nil)
				result <- err
			}()
			cancel()
			if err, _ := next(t, result); !errors.Is(err, context.Canceled) {
				t.Fatalf("got %v; want canceled", err)
			}
		})
	}
	deadline, stop := context.WithDeadline(ctx, time.Now().Add(-time.Second))
	defer stop()
	if _, _, err := Receive(deadline, nil); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v; want deadline exceeded", err)
	}
}

func TestE04Generate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assertInts(t, collect(t, Generate(ctx, 1, 0, -1)), []int{1, 0, -1})
	assertInts(t, collect(t, Generate(ctx)), nil)
	out := Generate(ctx, 1, 2, 3)
	cancel()        // 此前未消费，发送者也必须能退出。
	collect(t, out) // 取消与接收并发时，不强求“一个值都不能再返回”。
}

func TestE04GenerateCopiesInput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	values := []int{1, 2, 3}
	out := Generate(ctx, values...)
	for i := range values {
		values[i] = 99
	}
	assertInts(t, collect(t, out), []int{1, 2, 3})
}

func TestE05Map(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assertInts(t, collect(t, Map(ctx, source(1, 2, 3), func(n int) int {
		return n * n
	})), []int{1, 4, 9})
	out := Map(ctx, nil, func(n int) int { return n })
	cancel()
	if values := collect(t, out); len(values) != 0 {
		t.Fatal(values)
	}
}

func TestE05MapCancelDuringSend(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	computed := make(chan struct{})
	out := Map(ctx, source(1), func(n int) int {
		close(computed)
		return n
	})
	next(t, computed)
	cancel()
	collect(t, out)
}

func TestE06Merge(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	got := collect(t, Merge(ctx, source(1, 3), source(2, 4), source()))
	sort.Ints(got)
	assertInts(t, got, []int{1, 2, 3, 4})
	assertInts(t, collect(t, Merge(ctx)), nil)
	out := Merge(ctx, nil, nil)
	cancel()
	assertInts(t, collect(t, out), nil)
}

func TestE06MergePreservesEachInputOrder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	got := collect(t, Merge(ctx, source(1, 3, 5), source(2, 4, 6)))
	var odds, evens []int
	for _, n := range got {
		if n%2 == 1 {
			odds = append(odds, n)
		} else {
			evens = append(evens, n)
		}
	}
	assertInts(t, odds, []int{1, 3, 5})
	assertInts(t, evens, []int{2, 4, 6})
}

func TestE07Take(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assertInts(t, collect(t, Take(ctx, source(1, 2, 3), 2)), []int{1, 2})
	assertInts(t, collect(t, Take(ctx, source(1), 9)), []int{1})
	assertInts(t, collect(t, Take(ctx, nil, 0)), nil)
	assertInts(t, collect(t, Take(ctx, nil, -1)), nil)
	out := Take(ctx, nil, 1)
	cancel()
	assertInts(t, collect(t, out), nil)
}

func TestE08ParallelMap(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bad := errors.New("bad input")
	out, err := ParallelMap(ctx, source(1, 2, 3, 4), 3,
		func(_ context.Context, n int) (int, error) {
			if n == 3 {
				return 0, bad
			}
			return n * n, nil
		})
	if err != nil {
		t.Fatal(err)
	}
	values := collect(t, out)
	if len(values) != 4 {
		t.Fatalf("丢失或重复结果: %v", values)
	}
	seen := make(map[int]bool)
	for _, r := range values {
		if seen[r.Input] || r.Input < 1 || r.Input > 4 {
			t.Fatalf("异常输入: %+v", r)
		}
		seen[r.Input] = true
		if r.Input == 3 {
			if !errors.Is(r.Err, bad) {
				t.Fatalf("期望任务错误: %+v", r)
			}
		} else if r.Err != nil || r.Value != r.Input*r.Input {
			t.Fatalf("错误结果: %+v", r)
		}
	}
	for _, n := range []int{0, -1} {
		out, err := ParallelMap(ctx, nil, n, nil)
		if out != nil || !errors.Is(err, ErrInvalidWorkers) {
			t.Fatalf("非法并发数 %d: %v %v", n, out, err)
		}
	}
}

func TestE08ParallelMapBoundAndConcurrency(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const workers = 3
	started := make(chan struct{}, 30)
	gate := make(chan struct{})
	var active, peak atomic.Int32
	var once sync.Once
	defer once.Do(func() { close(gate) })
	fn := func(ctx context.Context, n int) (int, error) {
		current := active.Add(1)
		defer active.Add(-1)
		for old := peak.Load(); current > old; old = peak.Load() {
			if peak.CompareAndSwap(old, current) {
				break
			}
		}
		started <- struct{}{}
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-gate:
			return n * 2, nil
		}
	}
	out, err := ParallelMap(ctx, source(1, 2, 3, 4, 5, 6, 7), workers, fn)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < workers; i++ {
		next(t, started) // 三个任务必须能同时到达 gate，串行实现无法通过。
	}
	once.Do(func() { close(gate) })
	if got := collect(t, out); len(got) != 7 {
		t.Fatalf("got %d results", len(got))
	}
	if p := peak.Load(); p != workers {
		t.Fatalf("peak=%d; want %d", p, workers)
	}
	if active.Load() != 0 {
		t.Fatal("输出已关闭，但仍有回调执行")
	}
}

func TestE08ParallelMapCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{}, 1)
	finished := make(chan struct{})
	out, err := ParallelMap(ctx, source(1), 1,
		func(ctx context.Context, n int) (int, error) {
			defer close(finished)
			started <- struct{}{}
			<-ctx.Done()
			return 0, ctx.Err()
		})
	if err != nil {
		t.Fatal(err)
	}
	next(t, started)
	cancel()
	collect(t, out)
	next(t, finished)
}

func TestE09FirstSuccess(t *testing.T) {
	ctx := context.Background()
	errA, errB := errors.New("A"), errors.New("B")
	v, err := FirstSuccess(ctx,
		func(context.Context) (int, error) { return 0, errA },
		func(context.Context) (int, error) { return 42, nil },
	)
	if v != 42 || err != nil {
		t.Fatalf("got %d, %v", v, err)
	}
	_, err = FirstSuccess(ctx,
		func(context.Context) (int, error) { return 0, errA },
		func(context.Context) (int, error) { return 0, errB },
	)
	if !errors.Is(err, errA) || !errors.Is(err, errB) {
		t.Fatalf("必须保留全部错误: %v", err)
	}
	if _, err := FirstSuccess(ctx); !errors.Is(err, ErrNoJobs) {
		t.Fatal(err)
	}
}

func TestE09FirstSuccessCancelsLoser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started, exited := make(chan struct{}), make(chan struct{})
	result := make(chan error, 1)
	go func() {
		value, err := FirstSuccess(ctx,
			func(ctx context.Context) (int, error) {
				close(started)
				defer close(exited)
				<-ctx.Done()
				return 0, ctx.Err()
			},
			func(context.Context) (int, error) {
				<-started
				return 9, nil
			},
		)
		if err == nil && value != 9 {
			err = errors.New("wrong successful value")
		}
		result <- err
	}()
	if err, _ := next(t, result); err != nil {
		t.Fatal(err)
	}
	next(t, exited)
}

func TestE09FirstSuccessParentCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started, exited := make(chan struct{}), make(chan struct{})
	result := make(chan error, 1)
	go func() {
		_, err := FirstSuccess(ctx, func(ctx context.Context) (int, error) {
			close(started)
			defer close(exited)
			<-ctx.Done()
			return 0, ctx.Err()
		})
		result <- err
	}()
	next(t, started)
	cancel()
	if err, _ := next(t, result); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	next(t, exited)
}

func TestE10BatchSizeAndTail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out, err := Batch(ctx, source(1, 2, 3, 4, 5), 2, nil)
	if err != nil {
		t.Fatal(err)
	}
	got := collect(t, out)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v（检查底层数组是否被复用）", got, want)
	}
	for _, size := range []int{0, -1} {
		out, err := Batch(ctx, nil, size, nil)
		if out != nil || !errors.Is(err, ErrInvalidBatchSize) {
			t.Fatalf("size=%d: %v %v", size, out, err)
		}
	}
}

func TestE10BatchTickAndClosedTicks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	in := make(chan int)
	ticks := make(chan time.Time)
	out, err := Batch(ctx, in, 3, ticks)
	if err != nil {
		t.Fatal(err)
	}
	producerExited := make(chan struct{})
	go func() {
		defer close(producerExited)
		defer close(in)
		select {
		case in <- 1:
		case <-ctx.Done():
			return
		}
		select {
		case ticks <- time.Time{}:
		case <-ctx.Done():
			return
		}
		close(ticks)
		for _, n := range []int{2, 3} {
			select {
			case in <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	got := collect(t, out)
	if !reflect.DeepEqual(got, [][]int{{1}, {2, 3}}) {
		t.Fatalf("got %v", got)
	}
	next(t, producerExited)
}

func TestE10BatchCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out, err := Batch(ctx, nil, 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	cancel()
	if got := collect(t, out); len(got) != 0 {
		t.Fatal(got)
	}
}
