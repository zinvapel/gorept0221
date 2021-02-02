package main

import "fmt"

var i int64 = 10

func main() {
	fmt.Printf("'i' is %d, pointer of 'i' is %p\n", i, &i)

	i := 15
	fmt.Printf("'i' is %d, pointer of 'i' is %p\n", i, &i)

	if true {
		i := 20
		fmt.Printf("'i' is %d, pointer of 'i' is %p\n", i, &i)
	}
}
