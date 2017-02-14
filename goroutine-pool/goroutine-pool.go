package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	maxChannel = 3
	maxRoutine = 20
)

var (
	goroutineChannel chan int
)

func download(url string, pch *chan int) {
	fmt.Println(url)
	//fmt.Println("len(ch) is", len(*pch))

	time.Sleep(time.Duration(rand.Intn(300)) * time.Microsecond)

	// Complete all the jobs and receive from the channel.
	<- *pch
}

// Receive first, then the channel is ready immediately before the jobs complete.
//func download(pch *chan int) {
//	select {
//	case number := <- *pch:
//		fmt.Println(fmt.Sprintf("%d", number))
//		time.Sleep(time.Second)
//	}
//}

// Create the goroutine pool.
// The pch direct to the channel.
// THe max is the max goroutine number.
func createRoutinePool(pch *chan int, max int, fn func(string, *chan int))  {
	if *pch == nil {
		c := make(chan int, max)
		pch = &c
	} else {
		fmt.Println("pch is not nil")
		return
	}
	count := 0
	for {
		// Send to the channel.
		// If the channel is full, block.
		*pch <- count
		//fmt.Println("The channel pointer is", pch)
		go fn(fmt.Sprintf("%d", count), pch)
		count++

		if count > (maxRoutine - 1) {
			break
		}
	}
	close(*pch)

	// Call len() to check the channel is empty
	for len(*pch) != 0 {
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	//fmt.Println(goroutineChannel)
	createRoutinePool(&goroutineChannel, maxChannel, download)
	fmt.Println("Complete main()")
}
