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

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(2 * time.Second)
	done <- true
	fmt.Println("Done doing WORK")

	time.Sleep(3 * time.Hour)
}