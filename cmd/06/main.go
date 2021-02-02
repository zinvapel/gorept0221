package main

import "fmt"

func main() {
	Do()
}

func Do() {
	mp := map[string]string{"key": "value"}

	Inverse(mp)

	fmt.Println(mp)
}

func Inverse(m map[string]string) {
	// Написать код между...
	for k, v := range m {
		delete(m, k)
		m[v] = k
	}
	// ...этими комментариями
}
