package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(getError())
	fmt.Println(getErrorPlease())
}

func getErrorPlease() (err error) {
	// Пишем код здесь
	return
}

func getError() (err error) {
	if i, err := getIntAndError(); err != nil {
		fmt.Printf("%d\n", i)
	}

	return
}

func getIntAndError() (int, error) {
	return 0, errors.New("error")
}
