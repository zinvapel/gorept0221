package eight

func GetQueues() []chan string {
	return []chan string{
		make(chan string),
		make(chan string),
		make(chan string),
	}
}