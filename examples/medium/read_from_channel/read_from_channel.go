package read_from_channel

import (
	"fmt"
)

func readFromChannel[T any](ch <-chan T) {
	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("No data in the channel")
	}
}

func main() {
	intChannel := make(chan int)

	readFromChannel[int](intChannel)

	intChannel <- 10
	readFromChannel[int](intChannel)
}
