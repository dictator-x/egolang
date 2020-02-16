package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	// for {
	// 	if n, ok := <-c; ok {
	// 		fmt.Printf("Worker %d receiver %d\n", id, n)
	// 	}
	// }

	for n := range c {
		fmt.Printf("Worker %d receiver %c\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	wg.Wait()

}

func main() {
	chanDemo()
}
