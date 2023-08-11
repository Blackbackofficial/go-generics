package merge_сhannels

import (
	"fmt"
)

func mergeChannels[T any](a, b <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case val, ok := <-a:
				if !ok {
					a = nil
				} else {
					out <- val
				}
			case val, ok := <-b:
				if !ok {
					b = nil
				} else {
					out <- val
				}
			}
			if a == nil && b == nil {
				break
			}
		}
	}()

	return out
}

func main() {
	ints1 := make(chan int, 2)
	ints2 := make(chan int, 2)

	ints1 <- 1
	ints1 <- 2
	close(ints1)

	ints2 <- 3
	ints2 <- 4
	close(ints2)

	for val := range mergeChannels[int](ints1, ints2) {
		fmt.Println(val)
	}
}
