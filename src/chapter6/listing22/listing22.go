package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func run(channel chan int) {
	runner := <- channel
	fmt.Printf("runner %d keep the running\n", runner)
	var newRunner int
	if runner < 4 {
		newRunner = runner + 1
		go run(channel)
	}

	time.Sleep(time.Millisecond * 100)

	if runner == 4 {
		fmt.Printf("the running is completed.")
		wg.Done()
		return
	}

	fmt.Printf("%d runner is exchange runner %d\n", newRunner, runner)
	channel <- newRunner
}

func main() {
	wg.Add(1)

	channel := make(chan int)

	go run(channel)

	channel <- 1

	wg.Wait()
}
