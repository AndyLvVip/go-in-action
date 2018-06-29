package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var (
	grnum = 4
	tasknum = 10
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().Unix())
}

func work(id int, ch chan string) {
	defer wg.Done()

	for {
		task, ok := <- ch

		if !ok {
			fmt.Printf("Worker: %d shutdown\n", id)
			return
		}

		fmt.Printf("Worker: %d is working for %s\n", id, task)

		sleep := rand.Int63n(100)

		time.Sleep(time.Duration(sleep)  * time.Millisecond)

		fmt.Printf("Worker: %d complete %s\n", id, task)

	}
}

func main() {
	wg.Add(grnum)

	ch := make(chan string, tasknum)

	for task := 1; task <= tasknum; task++ {
		ch <- fmt.Sprintf("task %d", task)
	}

	for i := 1; i <= grnum; i++ {
		go work(i, ch)
	}


	close(ch)
	wg.Wait()
}
