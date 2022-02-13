package test

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "this is channel data"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second) // to make sure channel consumed
	close(channel)
}
