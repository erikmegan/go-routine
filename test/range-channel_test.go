package test

import (
	"fmt"
	"strconv"
	"testing"
)

// range channel will iterate recieve data until its done then close the channel
// used in case dont know how many data will recieved in channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "data ke: " + strconv.Itoa(i)
		}
		close(channel) // if not closed, it will deadlock because range channel will check continiously to recieved data until its closed
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("finish")
}
