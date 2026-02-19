package main

import (
	"errors"
	"fmt"
	"sync"
)

func workerWithErr(id int, wg *sync.WaitGroup, errs chan<- error) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work with a potential error outcome
	if id == 2 { // Example condition where a worker might fail
		errs <- errors.New(fmt.Sprintf("Worker %d encountered an error", id))
	}

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	errs := make(chan error, 3) // Capacity set to the number of goroutines

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerWithErr(i, &wg, errs)
	}

	wg.Wait()   // Ensure all workers finish
	close(errs) // Close the error channel

	for err := range errs {
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("All workers completed with error handling.")
}
