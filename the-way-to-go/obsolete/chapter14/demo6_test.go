package chapter14

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutines6(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump61(ch1)
	go pump62(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump61(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump62(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
