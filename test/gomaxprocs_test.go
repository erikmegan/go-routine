package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU: ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1) // above 0 will change total thread availability
	fmt.Println("total thread:", totalThread)

	totalGoRoutine := runtime.NumGoroutine() // will shown 2: 1 for run program, 1 for garbage collection deleting data from memory n go runtime
	fmt.Println("total goroutine:", totalGoRoutine)
}
