package main

import (
	"fmt"
	"os"
	"sync"
)

func readFile(file string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading", file, ":", err)
		return
	}
	ch <- fmt.Sprintln("File content:", string(data))

}

func main() {
	//using goroutines
	var wg sync.WaitGroup
	files := []string{"f1.txt", "f2.txt", "f3.txt"}
	ch := make(chan string)

	for _, file := range files {
		wg.Add(1)
		go readFile(file, &wg, ch)
	}

	for result := range ch {
		fmt.Println(result)
	}
	wg.Wait()
	close(ch)

}
// range ch waits until channel is closed

// But wg.Wait() is never reached

// So channel never closes

// So range waits forever

// 💥 Deadlock

// files := []string{"file1.txt", "file2.json", "file3.yaml"}

// for _, file := range files {
// 	file, err := os.ReadFile(file)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	fmt.Println("File content:", string(file))
// }

// file, err := os.ReadFile("file3.yaml")
// if err != nil {
// 	fmt.Println("Error reading file:", err)
// 	return
// }
// fmt.Println("File content:", string(file))
