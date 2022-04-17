package teller

import (
	"fmt"
	"sync"
	"time"
)

var depositsChannel = make(chan int) // send amount to deposit

var balancesChannel = make(chan int) // receive balance

// Write in the depositsChannel
func WriteDeposit(amount int) { depositsChannel <- amount }

// Read from balancesChannel
func ReadBalance() int { return <-balancesChannel }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		// read from [amount] from depositsChannel
		case amount := <-depositsChannel:
			// update balance
			balance += amount
		// update balancesChannel with balance
		case balancesChannel <- balance:
		}
	}
}

func Teller() {
	go teller() // start the monitor goroutine

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			WriteDeposit(1)
			time.Sleep(2 * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		for range ticker.C {
			balance := ReadBalance()
			fmt.Println(balance)
		}
	}()

	wg.Wait()
}
