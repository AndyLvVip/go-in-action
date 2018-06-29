package main

import (
	"sync"
	"fmt"
	"runtime"
)

func printPrime(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	next:
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}

	fmt.Println("Complete", prefix)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go printPrime("A", &wg)
	go printPrime("B", &wg)

	fmt.Println("waiting done")
	wg.Wait()
	fmt.Println("teminated")

}
