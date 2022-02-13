package test

import (
	"fmt"
	"testing"
)

func sendToChannel(channel chan string, data string) {
	channel <- data
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	go sendToChannel(channel1, "channel 1")
	defer close(channel1)

	channel2 := make(chan string)
	go sendToChannel(channel2, "channel 2")
	defer close(channel2)

	counter := 0 // total data from many channel
	for {        // infinity loop
		select {
		case data := <-channel1: // if data recieved from channel 1
			fmt.Println("data from channel1: ", data)
			counter++
		case data := <-channel2: // if data recieved from channel 2
			fmt.Println("data from channel2: ", data)
			counter++
		default:
			fmt.Println("waiting for data ...")
		}
		if counter == 2 {
			break
		}
	}
}
