package main

import "fmt"

func main() {
	// 演習1
	a := 10
	p := &a

	// 演習2
	fmt.Println("Value:", *p)

	// 演習3
	*p = 50
	fmt.Println("Updated a:", a)

	// 演習4
	result := isEven(&a)
	fmt.Println("Is Even:", result)
}

func isEven(n *int) bool {
	return *n%2 == 0
}
