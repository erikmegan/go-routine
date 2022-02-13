package test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld() // paralel and asyncronous
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("number: ", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
