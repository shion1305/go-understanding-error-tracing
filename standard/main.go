package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func doSomething() error {
	if err := doSomething1(); err != nil {
		return fmt.Errorf("new error: %w", err)
	}
	return nil
}

func doSomething1() error {
	return errors.New("test error")
}
