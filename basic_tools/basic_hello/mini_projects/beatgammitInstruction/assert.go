package main

import "fmt"

type NotBool struct {
	i interface{}
}

func (e *NotBool) Error() string {
	return fmt.Sprintf("Expected bool, but found: %T", e.i)
}

func NewNotBoolError(i interface{}) *NotBool {
	err := new(NotBool)
	err.i = i
	return err
}

func assertBool(i interface{}) (bool, error) {
	if b, ok := i.(bool); ok {
		return b, nil
	}
	return false, NewNotBoolError(i)
}

func main() {
	if b, err := assertBool(5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("It was a bool:", b)
	}

	if b, err := assertBool(true); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("It was a bool:", b)
	}
}
