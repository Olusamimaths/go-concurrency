package main

import (
	"fmt"
	"time"
)

func doSomething(x string) {
	fmt.Println(x)
}

func main_() {
	go doSomething("1")
	go doSomething("2")
	go doSomething("3")

	time.Sleep(2 * time.Second)

	fmt.Println("Inside main")
}