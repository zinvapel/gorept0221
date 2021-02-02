package main

import "fmt"

const (
	ExecuteRule = 1 << iota
	WriteRule
	ReadRule
)

func main() {
	fmt.Printf("%b\n", ReadRule | WriteRule | ExecuteRule)
}