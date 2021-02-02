package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	if _, ok := <-ch; ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed", <-ch)
	}
}