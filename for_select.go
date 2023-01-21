package main

import "fmt"

func main() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char:
			
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}