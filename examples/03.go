package main

import "fmt"

func main() {
	s := make([]int, 0, 0)
	m := make(map[int]int)
	c := make(chan int)
	fmt.Printf("%v %v %v\n", s, m, c)
	fmt.Printf("%T %T %T\n", s, m, c)

	sp := new([]int)
	sm := new(map[int]int)
	sc := new(chan int)
	fmt.Printf("%v %v %v\n", *sp, *sm, *sc)
	fmt.Printf("%T %T %T\n", sp, sm, sc)

	sp = &s
	sm = &m
	sc = &c
	fmt.Printf("%v %v %v\n", *sp, *sm, *sc)
	fmt.Printf("%T %T %T\n", sp, sm, sc)
}
