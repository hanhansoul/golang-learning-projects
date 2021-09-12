package chapter14

import (
	"fmt"
	"testing"
)

func TestGoroutines1(t *testing.T) {
	ch1 := make(chan int)
	go pump1(ch1)      // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}