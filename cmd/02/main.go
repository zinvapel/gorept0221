package main

import (
	"github.com/zinvapel/gorept0221/pkg/two"
	"log"
)

const (
	// Написать код между...
	ExecuteRule = 0
	WriteRule
	ReadRule
	// ...этими комментариями
)

func main() {
	Do()
}

func Do() {
	fileRole := two.GetFilePermission()

	if fileRole & ReadRule != 0 {
		log.Printf("%s\n", "Can read")
	}

	if fileRole & WriteRule != 0 {
		log.Printf("%s\n", "Can write")
	}

	if fileRole & ExecuteRule != 0 {
		log.Printf("%s\n", "Can execute")
	}
}