package main

import "fmt"

type Shape struct {
	X, Y float64
}

// Pointer receiver method: doubles X.
func (s *Shape) Double() {
	s.X *= 2
}

// Value receiver method: prints X.
func (s Shape) PrintX() {
	fmt.Println(s.X)
}

// Regular function that requires a pointer.
func DoubleX(s *Shape) {
	s.X *= 2
}

func main() {
	v := Shape{3, 4}
	p := &v

	// Pointer receiver methods can be called on values.
	v.Double()
	(&v).Double()

	// Value receiver methods can be called on pointers.
	p.PrintX()
	(*p).PrintX()

	//DoubleX(v) // compile error
	DoubleX(&v)
}
