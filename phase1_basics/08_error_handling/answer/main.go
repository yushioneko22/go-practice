package main

import (
	"errors"
	"fmt"
)

func Validate(s string) error {
	if s == "" {
		return errors.New("input is empty")
	}
	return nil
}

func main() {
	// 演習1
	err := errors.New("Something went wrong")
	fmt.Println(err)

	// 演習3
	errVal := Validate("")
	if errVal != nil {
		fmt.Println("Error:", errVal)
	}

	// 演習4
	wrapped := fmt.Errorf("Wrapped Error: %w", errVal)
	fmt.Println(wrapped)
}
