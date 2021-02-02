package main

import (
	"fmt"
	"github.com/zinvapel/gorept0221/pkg/eight"
	"sync"
)

func main() {
	Do()
}

func Do() {
	queues := eight.GetQueues()

	wg := &sync.WaitGroup{}

	for i, queue := range queues {
		wg.Add(1)

		// Написать код между...
		SendMessage(queue, i)
		// ...этими комментариями
	}

	go func() {
		for _, q := range queues {
			<-q
		}
	}()

	wg.Wait()
}

func SendMessage(queue chan string, i int) {
	queue <- fmt.Sprintf("msg for queue %d", i)
}
