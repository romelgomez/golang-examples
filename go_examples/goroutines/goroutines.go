package goroutines

import (
	"fmt"
	"sync"
)

func Goroutines() {

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go printNumber(i, &wg)
	}

	wg.Wait()
}

func printNumber(i int, wg *sync.WaitGroup) {
	fmt.Println(i)

	wg.Done()
}
