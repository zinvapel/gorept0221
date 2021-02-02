package main

import "fmt"

type MessageProvider struct {
	Queue chan<- string
}

func (p *MessageProvider) SendMessage(msg string) {
	go func() {
		p.Queue <- msg
	}()
}

type Mailbox struct {
	Name string
	Queue <-chan string
}

func (m *Mailbox) ReadMessage() string {
	return <-m.Queue
}

func main() {
	Do()
}

func Do() {
	// Пишем код здесь
	mp := MessageProvider{}
	mp.SendMessage("Hi, John")
	mp.SendMessage("I'm watching you")
	mp.SendMessage("Bye")

	personal := Mailbox{Name: "Personal"}

	for i := 0; i < 3; i++ {
		fmt.Println(personal.ReadMessage())
	}
}
