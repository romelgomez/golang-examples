package selects

import "fmt"

// if we have channels where we wanna read or write from, if any is blocking, we can use select to avoid that

func Selects() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	exitCh := make(chan interface{})

	go func() {
		ch1 <- 1
		// channel two is bloking
		// ch2 <- 1
		exitCh <- 1
	}()

	for {
		select {
		case i := <-ch1:
			fmt.Println("ch1: ", i)
		case i := <-ch2:
			fmt.Println("ch2: ", i)
		case i := <-exitCh:
			fmt.Println("exitCh: ", i)
			return
		}
	}

}
