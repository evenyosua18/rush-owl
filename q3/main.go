package main

import (
	"log"
	"time"
)

// worker model to define total worker and store task
type worker struct {
	TotalWorker int
	QueuedTaskC chan func()
}

func (w *worker) Run() {
	for i := 1; i <= w.TotalWorker; i++ {
		log.Printf("worker id %d: waiting for task\n", i)
		go func(id int) {
			for task := range w.QueuedTaskC {
				log.Printf("worker id %d: start processing task\n", id)
				task()
				log.Printf("worker id %d: finish processing task\n", id)
			}
		}(i)
	}
}

func (w *worker) AddTask(task func()) {
	w.QueuedTaskC <- task
}

// Result model to store result of calculation rectangle's area and to passed the data via channel
type Result struct {
	TaskID int
	Side   float64
	Area   float64
}

// CalculateRectangleArea function to calculate area of rectangle and give 3 seconds buffer
func CalculateRectangleArea(side float64) float64 {
	time.Sleep(3 * time.Second)
	return side * side
}

func main() {
	// initiate and run worker
	w := worker{TotalWorker: 3, QueuedTaskC: make(chan func())}
	w.Run()

	// save start time just for a record
	startTime := time.Now()

	// define total task
	totalTask := 9
	resultC := make(chan Result, totalTask)

	// assign task
	for i := 0; i < totalTask; i++ {
		taskID := i + 1
		side := float64(i) + 0.5
		w.AddTask(func() {
			log.Printf("add task %d\n", taskID)
			resultC <- Result{TaskID: taskID, Area: CalculateRectangleArea(side), Side: side}
		})
	}

	// get result of task
	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("task id %d and side %.1f finisihed with result %.2f", res.TaskID, res.Side, res.Area)
	}

	// just to record the total execution time
	log.Printf("finish all tasks with %f seconds\n", time.Now().Sub(startTime).Seconds())
}
