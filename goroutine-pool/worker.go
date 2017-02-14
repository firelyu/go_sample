package main

//http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
//http://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653369770&idx=1&sn=044be64c577a11a9a13447b373e80082

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	MaxWorker, _ = strconv.Atoi(os.Getenv("MAX_WORKERS"))
	MaxQueue, _  = strconv.Atoi(os.Getenv("MAX_QUEUES"))
)

// Payload
type Payload struct {
	Bill string
}

func (p Payload) UploadToS3() error {
	if p.Bill == "" {
		return errors.New("The Bill is missing.")
	}
	fmt.Println("Upload to S3", p.Bill)
	return nil
}

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job  // Store the Job for the Worker
	quit       chan bool // Quit the Worker
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		fmt.Println("Start the worker.")
		for {
			fmt.Println("Waiting...")
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					fmt.Println(err)
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		fmt.Println("Stop the work.")
		w.quit <- true
	}()
}

// Send job to JobQueue
func payloadHandler() {
	var content []Payload
	size := 20
	for i := 0; i < size; i++ {
		content = append(content, Payload{Bill: strconv.Itoa(i)})
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content {

		// let's create a job with the payload
		job := Job{Payload: payload}

		// Push the work onto the queue.
		JobQueue <- job
	}
}

// Dispatcher
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		//WorkQueue <- worker
		worker.Start()
	}

	go d.dispatch()
}

//func (d *Dispatcher) Stop()  {
//	close(JobQueue)
//	//for i := 0; i < len(WorkQueue); i++ {
//	//worker := <- WorkQueue
//	for worker := range WorkQueue {
//		worker.Stop()
//	}
//}

// Receive job from JobQueue
// Assign the Job to one Worker
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

//var WorkQueue chan Worker
func main() {
	fmt.Printf("The MAX_WORKERS is %d, the MAX_QUEUES is %d\n",
		MaxWorker, MaxQueue)
	JobQueue = make(chan Job, MaxQueue)
	//WorkQueue = make(chan Worker, MaxWorker)

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	payloadHandler()

	//for  {
	//	time.Sleep(100 * time.Millisecond)
	//}
	time.Sleep(2 * time.Second)

	//dispatcher.Stop()

	fmt.Println("main complete")
}
