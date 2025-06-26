package main

import (
	"errors"
	"fmt"
)

func someFunc() error {
	return errors.New("testy test test")
}

func main() {
	var err error

	// assign error, no checks
	err = someFunc()

	// Getting another error from closure, redeclaring err, no checks
	x, err := func() (int, error) {
		// reassigning error, no checks
		err = someFunc()
		return 0, err
	}()
	fmt.Println(x)
}

