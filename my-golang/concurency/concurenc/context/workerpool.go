package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost        int
}

func ticketProcessor(request <-chan ticketRequest, resul chan<- int) {
	for req := range request {
		fmt.Printf("Processing %d Tickes of PersonId %d wih cosT  %d",req.numTickets,req.personID,req.cost)
		time.Sleep(time.Second)
		resul <- req.personID
	}

}

func main() {
	numRequests := 5
	price := 5

	ticketRequests := make(chan ticketRequest,numRequests )
	ticketResults := make(chan int)

	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{
			personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price,
		}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}


}


// func worker(id int, tasks <-chan int, resuls chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d Processed task %d \n", id, task)
// 		time.Sleep(2 * time.Second)
// 		resuls <- task*2

// 	}

// }

// // time.Sleep(2*time.Second) // receiver := <-ch

// func main() {
// 	numofWorker := 3
// 	numOfJobs := 10

// 	tasks := make(chan int, numOfJobs)
// 	resul := make(chan int, numOfJobs)

// 	//creae worker
// 	for i := range numofWorker {
// 		go worker(i, tasks, resul)
// 	}

// 	//send values to task channel
// 	for i := range numOfJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	//collect the result
// 	for range numOfJobs {
// 		res := <-resul
// 		fmt.Println("Result", res)
// 	}

// }
