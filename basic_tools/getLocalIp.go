// getLocalIP.go file for learning purpose.
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Awesome Eric!")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("net.InterfaceAddrs -> %v", err)
	}

	for i, addr := range addrs {
		fmt.Println(i, addr)
	}
}
