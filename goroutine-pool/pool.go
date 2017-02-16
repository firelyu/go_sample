package main

// https://gobyexample.com/worker-pools

import (
	"fmt"
	"time"
)

type Job struct {
	JobId int
}

func worker(id int, job chan *Job, result chan bool) {
	go func() {
		for {
			fmt.Println("Waiting for job...")
			select {
			case j := <-job:
				// The job channel is closed.
				if j == nil {
					//fmt.Println("The channel is close", j)
					return
				}
				fmt.Println("worker", id, "started  job", j.JobId)
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finished job", j.JobId)
				result <- true
			}
		}
	}()
}

const (
	workerNumber        = 5
	jobNumber           = 10
	jobsChannelCache    = 100
	resultsChannelCache = 100
)

func main() {
	// According to the following link, it is better to set the job channel as pointer channel.
	// So after close the channel. The j:= <jobs get nil ptr.
	// http://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct/#answer-28449449
	jobs := make(chan *Job, jobsChannelCache)
	results := make(chan bool, resultsChannelCache)

	fmt.Println("len(jobs)", len(jobs))

	// Star worker goroutines.
	for i := 0; i < workerNumber; i++ {
		worker(i, jobs, results)
	}

	// Send job to the channel.
	time.Sleep(time.Second)
	for j := 0; j < jobNumber; j++ {
		jobs <- &Job{JobId:j}
	}
	close(jobs)

	for len(jobs) != 0 || len(results) != jobNumber{
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("len(jobs) is %d, len(results) is %d\n", len(jobs), len(results))

	fmt.Println("Complete main.")

}
