package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- 1
	}()

Loop:
	for {
		select {
		case <-ch1:
			fmt.Println("First channel")
			ch3 = nil
		case <-ch2:
			fmt.Println("Second channel")
		case <-ch3:
			fmt.Println("Third channel")
		case <- time.After(3 * time.Second):
			break Loop
		}
	}

	fmt.Printf("%p\n", ch3)
}