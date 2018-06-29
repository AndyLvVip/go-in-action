package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	wg sync.WaitGroup
)

func play(player string, court chan int) {

	defer wg.Done()

	for {
		ball, ok := <- court

		if !ok {
			fmt.Printf("%s you win\n", player)
			return
		}

		n := rand.Intn(100)

		if n % 13 == 0 {
			fmt.Printf("%s you lost\n", player)
			close(court)
			return
		}

		ball ++
		fmt.Printf("%s you hit is: %d\n", player, ball)
		court <- ball

	}

}

func main() {
	wg.Add(2)
	court := make(chan int)
	go play("andy", court)
	go play("jack", court)
	court <- 1
	wg.Wait()
}