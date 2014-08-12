package buffered

import (
	"time"
)

const workerCount = 4 // define the worker count in here.

var channel = make(chan []byte, 1024)
var workers = make(chan []*Worker, workerCount) // Cuz specific worker

func init() {
	for i := 0; i < workerCount; i++ {
		workers[i] = NewWorker(i)
		go workers[i].Work(channel)
	}
}

func Log(event []byte) {
	select {
	case channel <- event:
	case <-time.After(time.Second * 5): //throw away the msg, so sad, throw the msg
	}
}
