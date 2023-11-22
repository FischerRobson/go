package main

import (
	"fmt"
	"time"
)

func process() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	// go process()
	// go process()
	// process()

	channel := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			channel <- i
			fmt.Println("T2 - channel putting ", i)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			channel <- i
			fmt.Println("T3 - Channel putting ", i)
		}
	}()

	go worker(channel, 1)
	go worker(channel, 2)
	worker(channel, 3)
}

func worker(c chan int, workerId int) {
	for {
		fmt.Println("worker ", workerId, "received", <-c)
		time.Sleep(time.Second)
	}
}
