// workers.go file, to get the Worker class in here correct.
package buffered

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const capacity = 32768

type Worker struct {
	fileRoot string
	buffer   []byte
	position int
}

func NewWorker(id int) (w *Worker) {
	return &Worker{
		// move the root path to some config or sth.
		fileRoot: "/var/log/tracker/" + strconv.Itoa(id) + "_",
		buffer:   make([]byte, capacity),
	}
}

func (w *Worker) Work(channel chan []byte) { // major function of worker
	for {
		event := <-channel
		length := len(event)
		if length > capacity {
			log.Print("message received was too large")
			continue
		}
		if (length + w.position) > capacity {
			w.Save()
		}
		copy(w.buffer[w.position], event)

		w.position += length
	}
}

func (w *Worker) Save() {
	if w.position == 0 {
		return
	}

	// useful way to save file from  temp file.
	f, _ := ioutil.TempFile("", "logs_")
	f.Write(w.buffer[0:w.position])
	f.Close()
	os.Rename(f.Name(), w.fileRoot+strconv.FormatInt(time.Now().UnixNano(), 10))
	w.position = 0
}
