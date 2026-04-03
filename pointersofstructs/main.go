package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 2}
	p := &v
	fmt.Println("アドレスと中身を表示：", p)
	fmt.Print("アドレスを表示： ")
	fmt.Printf("%p\n", p)
	p.X = 324
	p.Y = 24
	fmt.Println(v)
}
