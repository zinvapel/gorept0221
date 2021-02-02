package main

import "fmt"

func main() {
	Do()
}

func Do() {
	slice := []int{1, 2, 3, 4, 5}

	multipliedSlice := Map(slice, Multiply(2))

	if slice[0] != multipliedSlice[0] {
		fmt.Println("Success")
	}
}

func Map(slice []int, apply func(int) int) []int {
	// Написать код между...
	for k := range slice {
		slice[k] = apply(slice[k])
	}

	return slice
	// ...этими комментариями
}

func Multiply(by int) func(int) int {
	return func(i int) int {
		return i * by
	}
}
