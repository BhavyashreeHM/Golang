package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at", tick)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}
}

// func periodicTask() {
// 	fmt.Println("Performing Periodic task at", time.Now())
// }

// func main() {
// 	timer := time.NewTicker(2 * time.Second)

// 	for {
// 		select {
// 		case <-timer.C:
// 			periodicTask()
// 		}

// 	}

// }

// func main() {
// 	timer := time.NewTicker(1*time.Second)
// 	defer timer.Stop()

// 	i:=1
// 	for range 5 {
// 		i++
// 		fmt.Println(i)
// 	}
// 	for range timer.C {
// 		i++
// 		fmt.Println(i)
// 	}
// }
