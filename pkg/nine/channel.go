package nine

import "time"

func GetFirst() chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- "First"
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}

func GetSecond() chan string {
	ch := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch <- "Second"
		}
	}()

	return ch
}

func GetThird() chan string {
	ch := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch <- "Third"
		}
	}()

	return ch
}
