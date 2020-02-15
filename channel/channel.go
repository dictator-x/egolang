package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// for {
	// 	if n, ok := <-c; ok {
	// 		fmt.Printf("Worker %d receiver %d\n", id, n)
	// 	}
	// }

	for n := range c {
		fmt.Printf("Worker %d receiver %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
