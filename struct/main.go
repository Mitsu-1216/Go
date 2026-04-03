package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z int
}

func main() {
	v := Vertex{3, 2, 4}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)
	v.X = 4
	fmt.Println(v)
}
