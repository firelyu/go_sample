package main

// https://gobyexample.com/worker-pools

import (
	"fmt"
	"time"
)

func worker(id int, job chan int, result chan bool) {
	go func() {
		for {
			fmt.Println("Waiting for job...")
			select {
			case j := <-job:
				fmt.Println("worker", id, "started  job", j)
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finished job", j)
				result <- true
			}
		}
	}()
}

const (
	workerNumber        = 3
	jobNumber           = 10
	jobsChannelCache    = 100
	resultsChannelCache = 100
)

func main() {
	jobs := make(chan int, jobsChannelCache)
	results := make(chan bool, resultsChannelCache)

	fmt.Println("len(jobs)", len(jobs))

	for i := 0; i < workerNumber; i++ {
		worker(i, jobs, results)
	}

	time.Sleep(time.Second)
	for j := 0; j < jobNumber; j++ {
		jobs <- j
	}
	//close(jobs)

	for len(jobs) != 0 || len(results) != jobNumber {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("len(jobs)", len(jobs))
}
