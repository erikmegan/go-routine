package test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "this is data" // send data to channel
		fmt.Println("finish send data to channel")
	}()
	data := <-channel // recieve data from channel
	fmt.Println("finish recieve data from channel")
	// will panic if no 1 recieve (deadlock)
	// how to make sure data from channel consumed ?

	fmt.Println(data)
	time.Sleep(3 * time.Second) // test to make sure goroutine finished
}
