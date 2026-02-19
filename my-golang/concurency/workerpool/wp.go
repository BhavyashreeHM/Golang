package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// worker function that processes an integer
func worker(id int, filechannel <-chan string, results chan<- string) {
	for file := range filechannel {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading", file, ":", err)
			return
		}
		time.Sleep(2* time.Second)
		results <- fmt.Sprintln("File content:", string(data))
	}
}

func main() {
	// const  = 5
	files := []string{"f1.txt", "f2.txt", "f3.txt", "f4.txt", "f5.txt"}

	filechannel := make(chan string, len(files))
	results := make(chan string, len(files))

	var wg sync.WaitGroup

	// start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, filechannel, results)
		}(w)
	}

	// send filechannel
	for _, file := range files {
		filechannel <- file
	}
	close(filechannel)

	// wait for workers to finish
	wg.Wait()
	close(results)

	// collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
