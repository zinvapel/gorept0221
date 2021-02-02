package main

import "fmt"

func PrintAddr(i interface{}) {
	fmt.Printf("%v %p\n", i, &i)
}

func main() {
	i := 0
	fmt.Printf("%v %p\n", i, &i)
	PrintAddr(i)

	p := new(int)
	fmt.Printf("%v %p\n", p, &p)
	PrintAddr(p)

	m := make(map[string]string, 0)
	fmt.Printf("%v %p\n", m, &m)
	PrintAddr(m)

	s := make([]string, 0)
	fmt.Printf("%v %p\n", s, &s)
	PrintAddr(s)

	c := make(chan string)
	fmt.Printf("%v %p\n", c, &c)
	PrintAddr(c)
}

//map:
//struct {
//	m *internalHashtable
//}
//channel:
//struct {
//	c *internalChannel
//}
//slice:
//struct {
//	array *internalArray
//	len   int
//	cap   int
//}
//function:
//struct {
//	func *internalFunc
//}