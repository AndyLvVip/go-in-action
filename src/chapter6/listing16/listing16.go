package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wg sync.WaitGroup
	mx sync.Mutex
)

func incCounter(id int) {

	defer wg.Done()

	for count := 0; count < 2; count ++ {
		mx.Lock()
		value := counter
		runtime.Gosched()
		value ++
		counter = value
		mx.Unlock()
	}

}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()

	fmt.Printf("counter: %d\n", counter)
}
