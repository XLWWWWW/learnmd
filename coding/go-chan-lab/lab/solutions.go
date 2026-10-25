//go:build solution

package channellab

import (
	"context"
	"errors"
	"sync"
	"time"
)

func Sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}

func TrySend(out chan<- int, n int) bool {
	select {
	case out <- n:
		return true
	default:
		return false
	}
}

func Receive(ctx context.Context, in <-chan int) (int, bool, error) {
	if err := ctx.Err(); err != nil {
		return 0, false, err
	}
	select {
	case <-ctx.Done():
		return 0, false, ctx.Err()
	case n, ok := <-in:
		return n, ok, nil
	}
}

func Generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	// 复制切片，避免函数返回后调用方改动底层数组。
	values := append([]int(nil), nums...)
	go func() {
		defer close(out)
		for _, n := range values {
			if ctx.Err() != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func Map(ctx context.Context, in <-chan int, fn func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			if ctx.Err() != nil {
				return
			}
			var n int
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				n = v
			}
			// 先计算，再进入 select；fn 应是短时间内完成的纯计算。
			value := fn(n)
			select {
			case <-ctx.Done():
				return
			case out <- value:
			}
		}
	}()
	return out
}

func Merge(ctx context.Context, inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(in <-chan int) {
			defer wg.Done()
			for {
				if ctx.Err() != nil {
					return
				}
				var n int
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					n = v
				}
				select {
				case <-ctx.Done():
					return
				case out <- n:
				}
			}
		}(input)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Take(ctx context.Context, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			if ctx.Err() != nil {
				return
			}
			var value int
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				value = v
			}
			select {
			case <-ctx.Done():
				return
			case out <- value:
			}
		}
	}()
	return out
}

func ParallelMap(ctx context.Context, in <-chan int, workers int,
	fn func(context.Context, int) (int, error)) (<-chan Result, error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	out := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				if ctx.Err() != nil {
					return
				}
				var n int
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					n = v
				}
				value, err := fn(ctx, n)
				result := Result{Input: n, Value: value, Err: err}
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out, nil
}

func FirstSuccess(ctx context.Context,
	jobs ...func(context.Context) (int, error)) (int, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	if len(jobs) == 0 {
		return 0, ErrNoJobs
	}
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	// 每个 job 只发送一次。即使调用者已返回，也足以容纳全部结果。
	results := make(chan Result, len(jobs))
	for _, job := range jobs {
		go func(run func(context.Context) (int, error)) {
			value, err := run(child)
			results <- Result{Value: value, Err: err}
		}(job)
	}
	errs := make([]error, 0, len(jobs))
	for i := 0; i < len(jobs); i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case r := <-results:
			if r.Err == nil {
				return r.Value, nil
			}
			errs = append(errs, r.Err)
		}
	}
	return 0, errors.Join(errs...)
}

func Batch(ctx context.Context, in <-chan int, size int,
	ticks <-chan time.Time) (<-chan []int, error) {
	if size <= 0 {
		return nil, ErrInvalidBatchSize
	}
	out := make(chan []int)
	go func() {
		defer close(out)
		batch := make([]int, 0, size)
		flush := func() bool {
			if len(batch) == 0 {
				return true
			}
			select {
			case <-ctx.Done():
				return false
			case out <- batch:
				// 交出底层数组的所有权，后续批次不能复用它。
				batch = make([]int, 0, size)
				return true
			}
		}
		for {
			if ctx.Err() != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				if !ok {
					flush()
					return
				}
				batch = append(batch, n)
				if len(batch) == size && !flush() {
					return
				}
			case _, ok := <-ticks:
				if !ok {
					ticks = nil
					continue
				}
				if !flush() {
					return
				}
			}
		}
	}()
	return out, nil
}
