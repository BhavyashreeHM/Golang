package main

import (
	"fmt"
	"sync"
)

func wait(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Im from wait group")
}

func mai() {
	var wg sync.WaitGroup

	wg.Add(1)
	go wait(&wg)

	wg.Wait()
	fmt.Println("Its Done")
}
