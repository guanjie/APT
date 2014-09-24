// multiplexing.go file for learning purpose. NOT COMPLETE, NOT RUNNABLE
package main

import (
	"fmt"
)

type request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

// type Reply struct{}
//
// type Request struct {
// 	arg1, arg2 int
// 	replyc     chan *Reply
// }

func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service <-chan *request) { // send in request from service channel, then goroutine the run() function
	for {
		req := <-service
		go run(op, req)
	}
}

// func startServer(op binOp) chan<- *request {
// 	service := make(chan *request)
// 	go server(op, service)
// 	return service
// }

func startServer(op binOp) chan<- *request {
	service := make(chan *request)
	go server(op, req)
	return service
}

func main() {
	fmt.Println("Awesome Eric!")

	adderChan := startServer(
    func(a, b int) int { return a + b }
    )
}
