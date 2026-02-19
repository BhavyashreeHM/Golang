package main

import (
	"fmt"
	"sync"
	"time"
)

// ============ BASIC EXAMPLE WITHOUT USING CHANNELS
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// wg.Add(1)  /// WRONG PRACTICE
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // simulate some time spent on processing the task
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAIT GROUPS

	// Launch workers
	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("All workers finished")
}
