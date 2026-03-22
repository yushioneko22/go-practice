package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

type Describer interface {
	Describe() string
}

type User struct {
	Name string
}

func (u User) Describe() string {
	return "User name is " + u.Name
}

func printDesc(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	// 演習1
	c := &Counter{count: 0}
	c.Increment()
	fmt.Println("Count:", c.count)

	// 演習4
	u := User{Name: "Alice"}
	printDesc(u)
}
