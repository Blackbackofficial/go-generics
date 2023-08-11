package race_context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result[T any] struct {
	Value T
	Err   error
}

func RaceContext[T any](ctx context.Context, funcs ...func(context.Context) (T, error)) (Result[T], error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan Result[T], len(funcs))
	wg := sync.WaitGroup{}

	for _, fn := range funcs {
		wg.Add(1)
		go func(fn func(context.Context) (T, error)) {
			defer wg.Done()

			res, err := fn(ctx)
			if err != nil {
				ch <- Result[T]{Err: err}
				return
			}
			ch <- Result[T]{
				Value: res,
				Err:   nil,
			}
		}(fn)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case res := <-ch:

		return res, nil
	case <-ctx.Done():
		return Result[T]{}, ctx.Err()
	}
}

func main() {
	ctx := context.Background()
	funcs := []func(context.Context) (int, error){
		func(ctx context.Context) (int, error) {
			time.Sleep(2 * time.Second)
			return 1, nil
		},
		func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Second)
			return 2, nil
		},
	}

	output, err := RaceContext(ctx, funcs...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if output.Err != nil {
		fmt.Println("Function Error:", output.Err)
		return
	}
	fmt.Println("Result:", output.Value)
}
