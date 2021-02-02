package main

import "fmt"

func changeMap(m map[string]string) {
	m["key"] = "value"
}

func changeSlice(s []string) {
	s[0] = "value"
}

func changeChan(c chan string) {
	c <- "value"
}

func changeArr(a [5]string) {
	a[4] = "value"
}

func main() {
	m := make(map[string]string, 0)
	changeMap(m)
	fmt.Printf("%v\n", m)

	s := make([]string, 1)
	changeSlice(s)
	fmt.Printf("%v\n", s)

	c := make(chan string)
	go changeChan(c)
	fmt.Printf("%s\n", <-c)

	a := [5]string{"1", "2", "3", "4", "5"}
	changeArr(a)
	fmt.Printf("%s\n", a[4])
}