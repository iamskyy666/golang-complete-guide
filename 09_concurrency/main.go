package main

import (
	"fmt"
	"sync"
)

// 📂 09_concurrency
// Closing Channels - with WaitGroups

func main() {
	jobs:=make(chan int,5)
	//done:=make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)

	// start a goroutine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r,ok:=<-jobs
			if ok {
				fmt.Println("Got this message!",r)
			}else{
				//done<-false
				fmt.Println("Channel Closed..!")
				return
			}
		}
	}(&wg)

		for i := 1; i <= 3; i++ {
			jobs<-i
			fmt.Println("Sending..",i)
		}

		close(jobs)
	// <-done
	wg.Wait()
}

// $ go run main.go
// Sending.. 1
// Sending.. 2
// Sending.. 3
// Got this message! 1
// Got this message! 2
// Got this message! 3
// Channel Closed..!

