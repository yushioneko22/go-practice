package main

import "fmt"

type Rect struct {
	Width, Height float64
}

// 値レシーバ (読み取り専用)
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// ポインタレシーバ (更新用)
func (r *Rect) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

// インターフェースの定義
type Shaper interface {
	Area() float64
}

func printArea(s Shaper) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	r := Rect{Width: 10, Height: 5}
	fmt.Println("Initial Area:", r.Area())

	r.Scale(2)
	fmt.Println("Scaled Area:", r.Area())

	printArea(r) // Shaperインターフェースを満たしている
}
