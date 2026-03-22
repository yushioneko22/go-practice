package main

import "fmt"

func main() {
	// 演習1: 1から5まで
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 演習2: 合計
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 演習3: range
	fruits := []string{"apple", "orange", "grape"}
	for i, f := range fruits {
		fmt.Printf("%d: %s\n", i, f)
	}

	// 演習4: break
	a := 0
	for {
		if a == 5 {
			fmt.Println("Break at 5")
			break
		}
		a++
	}
}
