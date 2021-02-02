package main

import (
	"fmt"
	"time"
)

func main() {
	for k := range make([]int, 3) {
		fmt.Printf("Key is %d\n", k)

		go func() {
			fmt.Printf("Key is %d\n", k)
		}()
	}

	time.Sleep(time.Second)
}
