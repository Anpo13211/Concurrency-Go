package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task func()

// Worker represents a worker with a task deque
type Worker struct {
	id    int
	tasks chan Task
}

// NewWorker creates a new worker
func NewWorker(id int, taskQueue chan Task) *Worker {
	return &Worker{
		id:    id,
		tasks: taskQueue,
	}
}

// Start starts the worker
func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup, workers []*Worker) {
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case task := <-w.tasks:
				task()
			default:
				// Try to steal a task from another worker
				for _, worker := range workers {
					if worker != w {
						select {
						case task := <-worker.tasks:
							fmt.Printf("Worker %d stole a task from worker %d\n", w.id, worker.id)
							task()
						default:
							fmt.Println("No tasks to steal.")
							cancel()
						}
					}
				}
			}
		}
	}()
}

// AddTask adds a task to the worker's queue
func (w *Worker) AddTask(task Task) {
	w.tasks <- task
}

func main() {
	numWorkers := 4
	workers := make([]*Worker, numWorkers)
	taskQueues := make([]chan Task, numWorkers)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	// Initialize workers and their task queues
	for i := 0; i < numWorkers; i++ {
		taskQueues[i] = make(chan Task, 10) // Adjust buffer size as needed
		workers[i] = NewWorker(i, taskQueues[i])
	}

	// Start workers
	for _, worker := range workers {
		worker.Start(ctx, &wg, workers)
	}

	// Distribute tasks among workers
	for i := 0; i < 20; i++ {
		i := i // Create a copy of i for the closure
		worker := workers[i%numWorkers]
		worker.AddTask(func() {
			fmt.Printf("Processing task %d\n", i)
			time.Sleep(100 * time.Millisecond) // Simulate task processing time
		})
	}

	// Wait for all workers to finish
	wg.Wait()
}
