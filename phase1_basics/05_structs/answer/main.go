package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	// 演習2
	p := Product{Name: "Laptop", Price: 1000}

	// 演習3
	p.Price = 500
	fmt.Printf("Product: %s, Price: %d\n", p.Name, p.Price)

	// 演習4
	showProduct(p)
}

func showProduct(p Product) {
	fmt.Printf("[Name: %s, Price: %d]\n", p.Name, p.Price)
}
