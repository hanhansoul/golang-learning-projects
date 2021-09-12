package chapter14

import (
	"fmt"
	"testing"
	"time"
)

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
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

func TestExample41(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func tel1(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func tel() (<-chan int, <-chan interface{}) {
	done := make(chan interface{})
	ch := make(chan int)
	go func() {
		defer close(done)
		defer close(ch)
		for i := 0; i < 15; i++ {
			ch <- i
		}
	}()
	return ch, done
}

func TestExample42(t *testing.T) {
	var ok = true
	ch, done := tel()
	for ok {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Printf("The counter is at %d\n", i)
			}
		case <-done:
			return
		}
	}
}

func TestExample43(t *testing.T) {

}
