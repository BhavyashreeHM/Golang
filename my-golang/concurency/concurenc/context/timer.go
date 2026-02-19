package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-timer1.C:
			fmt.Println("timer1 executed")
		case <-timer2.C:
			fmt.Println("Timer2 executed")

		}
	}
	// time.Sleep(2*time.Second)
}

// func main(){
// 	timer := time.NewTimer(2 *time.Second) //non-blocking

// 	go func(){
// 		<-timer.C
// 		fmt.Println("Delayed Operation executed")
// 	}()
// 	fmt.Println("Waiting for goroutine")
// 	time.Sleep(3*time.Second) //blocking
// 	fmt.Println("End of Program")

// }

//Delayed Operation
// func longRunningOperation(){
// 	for i:= range 20{
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeOut := time.After(2*time.Second)
// 	done := make(chan bool)

// 	go func (){
// 		longRunningOperation()
// 		done<-true
// 	}()

// 	select {
// 	case <-timeOut:
// 		fmt.Println("Operation time out")
// 	case <-done:
// 		fmt.Println("Operation Completed")

// 	}
// }

// func main() {
// 	// time.Sleep(2*time.Second) //blocking by nature
// 	fmt.Println("Starting App")
// 	timer := time.NewTimer(5 * time.Second) // not block but it will create timer of 2seconds

// 	stoped:= timer.Stop()
// 	if stoped{
// 		fmt.Println("Stopped")
// 	}
// 	go func() {
// 		fmt.Println("Timer Expired")
// 	}()
// 	fmt.Println("Waiting for time.c")
// 	timer.Reset(time.Second)
// 	fmt.Println("Time Reset")
// 	<-timer.C //block and wait for that second specified  in timer
// }
