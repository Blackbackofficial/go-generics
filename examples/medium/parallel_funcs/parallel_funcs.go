package parallel_funcs

import (
	"fmt"
	"sync"
)

func parallelFuncs[T any](funcs []func() T) []T {
	var wg sync.WaitGroup
	results := make([]T, len(funcs))

	for i, function := range funcs {
		wg.Add(1)
		go func(i int, function func() T) {
			defer wg.Done()
			results[i] = function()
		}(i, function)
	}

	wg.Wait()

	return results
}

func main() {
	intFuncs := []func() int{
		func() int { return 1 },
		func() int { return 2 },
	}

	stringFuncs := []func() string{
		func() string { return "Hello" },
		func() string { return "World" },
	}

	resultsInts := parallelFuncs(intFuncs)
	resultsStrings := parallelFuncs(stringFuncs)
	fmt.Println(resultsInts, resultsStrings)
}
