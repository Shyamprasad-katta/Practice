package main

import (
	"fmt"
	"sync"
)

func Worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d", id, job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int)

	var wg sync.WaitGroup

	for j := 0; j < 3; j++ {
		wg.Add(1)
		go Worker(j, jobs, results, &wg)
	}
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	// close(results)
	go func() {
		wg.Wait()
		// close(results)
	}()
	for result := range results {
		fmt.Println("Result:", result)
	}
}
A












































