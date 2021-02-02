package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[:]

	slice2[0] = 9
	fmt.Println(slice)

	slice3 := make([]int, 0, 0)
	// Копирует меньшее (по len) из 2-х слайсов
	copy(slice3, slice)
	fmt.Println(slice3)

	slice4 := make([]int, len(slice), len(slice))
	copy(slice4, slice)
	slice4[0] = 10
	fmt.Println(slice4)
	fmt.Println(slice)
}
