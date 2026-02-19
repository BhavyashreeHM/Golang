package main

import "fmt"

type Task struct {
	ID   int
	Work func() error
}

type Worker struct {
	ID     int
	TaskCh chan Task
	DoneCh chan bool
}

// Start method to start the worker process
func (w Worker) Start() {
	go func() {
		for task := range w.TaskCh {
			err := task.Work()
			if err != nil {
				// handle error
			}
			w.DoneCh <- true
		}
	}()
}

func CreateWorkerPool(numWorkers int, tasks []Task) {
	taskCh := make(chan Task, len(tasks))
	doneCh := make(chan bool)

	var workers []Worker
	for i := 1; i <= numWorkers; i++ {
		worker := Worker{
			ID:     i,
			TaskCh: taskCh,
			DoneCh: doneCh,
		}
		worker.Start()
		workers = append(workers, worker)
	}

	// Enqueue tasks
	for _, task := range tasks {
		taskCh <- task
	}

	// Wait for all tasks to complete
	for i := 0; i < len(tasks); i++ {
		<-doneCh
	}
}

func main() {

	tasks := []Task{
		{1, func() error { fmt.Println("Task 1"); return nil }},
		{2, func() error { fmt.Println("Task 2"); return nil }},
		{3, func() error { fmt.Println("Task 3"); return nil }},
	}
	CreateWorkerPool(2, tasks)

}
