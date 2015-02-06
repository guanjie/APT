package main

import "fmt"

type Appender struct {
	data []byte
}

// Overwrite the Write method of Appender(Write method is a predefined method that's high level)
func (a *Appender) Write(p []byte) (int, error) {
	newArr := append(a.data, p...)
	a.data = newArr
	return len(p), nil
}

func main() {
	appender := new(Appender)
	appender.data = []byte{}

	fmt.Fprintf(appender, "Hello %s", "World")
	fmt.Fprintf(appender, "\nI just added %d %s", 2, "lines")

	fmt.Println(string(appender.data))
}
