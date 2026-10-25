package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	lab "example.com/go-chan-lab/lab"
)

func main() {
	scenario := flag.String("scenario", "basic", "basic|pipeline|workers|first|batch")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	switch *scenario {
	case "basic":
		ch := make(chan int, 2)
		ch <- 0
		ch <- 7
		close(ch)
		for i := 0; i < 3; i++ {
			value, ok := <-ch
			fmt.Printf("value=%d ok=%v\n", value, ok)
		}
	case "pipeline":
		in := lab.Generate(ctx, 1, 2, 3, 4)
		squares := lab.Map(ctx, in, func(n int) int { return n * n })
		for n := range lab.Take(ctx, squares, 2) {
			fmt.Println(n)
		}
		// 提前停止消费时，必须通知仍在工作的上游。
		cancel()
	case "workers":
		in := lab.Generate(ctx, 1, 2, 3, 4, 5)
		out, err := lab.ParallelMap(ctx, in, 3,
			func(ctx context.Context, n int) (int, error) {
				if n == 3 {
					return 0, errors.New("示例：任务 3 失败")
				}
				return n * n, nil
			})
		if err != nil {
			panic(err)
		}
		for r := range out {
			fmt.Printf("input=%d value=%d error=%v\n", r.Input, r.Value, r.Err)
		}
	case "first":
		deadline, stop := context.WithTimeout(ctx, time.Second)
		defer stop()
		value, err := lab.FirstSuccess(deadline,
			func(context.Context) (int, error) { return 0, errors.New("副本 A 失败") },
			func(context.Context) (int, error) { return 42, nil },
		)
		fmt.Println(value, err)
	case "batch":
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		in := lab.Generate(ctx, 1, 2, 3, 4, 5)
		out, err := lab.Batch(ctx, in, 2, ticker.C)
		if err != nil {
			panic(err)
		}
		for values := range out {
			fmt.Println(values)
		}
	default:
		panic("unknown scenario: " + *scenario)
	}
}
