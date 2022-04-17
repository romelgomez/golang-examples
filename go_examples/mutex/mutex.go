package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// go run -race main.go

func Mutex() {

	var state = make(map[int]int)

	var syncMutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)

				// To safe run in only one thread
				syncMutex.Lock()
				// this part of the code
				total += state[key]
				syncMutex.Unlock()

				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				syncMutex.Lock()
				state[key] = val
				syncMutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops:", writeOpsFinal)

	syncMutex.Lock()
	fmt.Println("state", state)
	syncMutex.Unlock()

}
