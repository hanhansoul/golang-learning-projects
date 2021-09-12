package chapter14

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutines2(t *testing.T) {
	out := make(chan int, 2)
	fmt.Println("before")
	out <- 2
	fmt.Println("after")
	go f1(out)
	time.Sleep(1e9)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
