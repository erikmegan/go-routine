package test

import (
	"fmt"
	"testing"
)

func TestBufferedChanneL(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "first data" // only send data to channel and no 1 consume. if no buffer then will deadlock because goroutine is blockin waiting someone recieve data

	fmt.Println("selesai")
}
