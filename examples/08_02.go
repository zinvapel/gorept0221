package main

import (
	"fmt"
	"time"
)

func main() {
	for k := range make([]int, 3) {
		fmt.Printf("Key is %d\n", k)

		k := k
		go func() {
			fmt.Printf("Key is %d\n", k)
		}()
	}

	for k := range make([]int, 3) {
		fmt.Printf("Key is %d\n", k)

		go func(k int) {
			fmt.Printf("Key is %d\n", k)
		}(k)
	}

	time.Sleep(time.Second)
}
