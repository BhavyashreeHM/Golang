package main

import (
	"context"
	"fmt"

	"time"
)

func doWork(cx context.Context) {
	for {
		select {
		case <-cx.Done():
			fmt.Println("Work Cancelled:", cx.Err())
			return

		default:
			fmt.Println("Working........")

		}
		time.Sleep(500 * time.Second)
	}
}

// receiver := <-ch
func main() {
	cx := context.Background()
	cx, cancel := context.WithTimeout(cx, 2*time.Second)
	defer cancel()

	cx = context.WithValue(cx, "USERID", "ghjklsdkjhjd")
	go doWork(cx)

	time.Sleep(3 * time.Second)

	requesID := cx.Value("USERID")
	if requesID != nil {
		fmt.Println("Reques ID is", requesID)

	} else {
		fmt.Println("No Reques ID found")

	}

}
