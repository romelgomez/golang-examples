package channels

import (
	"fmt"
	"sync"
)

func Channels() {
	// simple()
	// blocking()
	// Buffered()
	// Nil()
	// BrokenFor()
	// Range()
	// Closing()
	Broadcast()
}

func blocking() {
	channel := make(chan bool)

	// writing
	go func() { channel <- true }()

	// reading
	fmt.Println("blocking channel", <-channel)
}

func simple() {
	channel := make(chan int)

	go func() { channel <- 1 }()

	a := <-channel

	fmt.Println("simple channel", a)
}

func Buffered() {
	// we have 1 slot
	channel := make(chan bool, 1)
	// writing
	channel <- true
	// reading
	fmt.Println(<-channel)

	// we have 2 slot
	channel2 := make(chan bool, 2)
	// writing
	channel2 <- false
	channel2 <- true
	// reading
	fmt.Println(<-channel2)
}

// will block forever
func Nil() {
	var nilChannel chan bool
	// reading
	<-nilChannel
}

// Will only print the 1, 2, the 3 nop, it ends before the loop completed
func BrokenFor() {
	channel := make(chan int)

	go func() {
		// infinite lop
		for {
			// reading
			fmt.Println(<-channel)
		}
	}()

	// Reading
	channel <- 1
	channel <- 2
	channel <- 3

}

// Will print all the values in the channel
func Range() {
	channel := make(chan int)

	go func() {
		// Writing
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
}

// RULE OF thumb: only the producer CLOSE the channel
func Closing() {
	// Closing a channels
	// 1. Wrinting to a closed channels will panic
	channel := make(chan int)
	close(channel)
	// channel <- 3 // panic

	// 2. Reading from a closed channel will always succeed (unless in range)
	channel2 := make(chan int)
	close(channel2)

	// the first value will be zero-value
	// Second is always a boolean indicating if the channels is open

	// Reading
	val, ok := <-channel2
	fmt.Printf("val is: %v, is open %v\n", val, ok)
}

func Broadcast() {
	var wg sync.WaitGroup
	channel := make(chan int)

	wg.Add(2)

	go func(c <-chan int) {
		for i := range c {
			fmt.Println("1st goroutine", i)
		}
		wg.Done()
	}(channel)

	go func(c <-chan int) {
		for i := range c {
			fmt.Println("2st goroutine", i)
		}
		wg.Done()
	}(channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)
	wg.Wait()
}
