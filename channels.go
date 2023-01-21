package main

import "fmt"

func main_2() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func ()  {
		anotherChannel <- "dog"
	}()

	var msg string
	select {
	case msg := <- myChannel:
		fmt.Println("Message from my channel, " + msg)
	case msg := <- anotherChannel:
		fmt.Println("Message from another channel, " + msg)
	}

	fmt.Println(msg)
}