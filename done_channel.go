package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing WORK")
		}
	}
}

func main_4() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(2 * time.Second)
	close(done)
	fmt.Println("Done doing WORK")

	time.Sleep(3 * time.Hour)
}