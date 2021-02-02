package main

import (
	"context"
	"fmt"
	"github.com/zinvapel/gorept0221/pkg/nine"
)


func main() {
	Do(context.Background())
}

func Do(ctx context.Context) {
	f, s, t := nine.GetFirst(), nine.GetSecond(), nine.GetThird()

	for {
		select {
		// Написать код между...
		case msg := <-f:
			fmt.Println(msg)
		case msg := <-s:
			fmt.Println(msg)
		case msg := <-t:
			fmt.Println(msg)
		// ...этими комментариями
		case <-ctx.Done():
			return
		}
	}
}