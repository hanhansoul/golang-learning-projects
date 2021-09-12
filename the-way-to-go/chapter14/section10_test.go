package chapter14

import (
	"fmt"
	"testing"
)

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request) {
	for {
		req := <-service // requests arrive here
		// start goroutine for request:
		go run(op, req) // don't wait for op
	}
}

func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}

func TestExample1001(t *testing.T) {
	adder := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	// checks:
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	fmt.Println("done")
}

func TestExample1002(t *testing.T) {
	const MAXREQS = 50

	var sem = make(chan int, MAXREQS)

	process := func(r *Request) {
		// do something
	}

	handle := func(r *Request) {
		sem <- 1 // doesn't matter what we put in it
		process(r)
		<-sem // one empty place in the buffer: the next request can start
	}

	server := func(service chan *Request) {
		for {
			request := <-service
			go handle(request)
		}
	}

	service := make(chan *Request)
	go server(service)
}
