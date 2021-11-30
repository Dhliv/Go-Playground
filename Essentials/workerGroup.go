package Essentials

import "fmt"

func fibonacci(n int) int64 {
	if n <= 1 {
		return int64(1)
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(id int, jobs <-chan int, results chan<- int64) {
	for job := range jobs {
		fmt.Printf("Worker with ID %d started fib with n = %d\n", id, job)
		fib := fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func WorkerGroup() {
	tasks := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 30, 32, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int64, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go worker(i, jobs, results) // we create workers that will be waiting tasks.
	}

	for _, value := range tasks {
		jobs <- value // We assign tasks for workers.
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results // We receive the results of the operations. In this case, we just throw it away.
	}
}
