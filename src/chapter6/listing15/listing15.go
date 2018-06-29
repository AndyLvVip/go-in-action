package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var (
	shutdown int64
	wg sync.WaitGroup
)


func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			break
		}
	}
}

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}
