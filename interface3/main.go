package main

import "fmt"

type Shaper interface {
	Area() int
}

type Shape struct {
	X, Y int
}

// ポインタレシーバでメソッドを定義
func (s *Shape) Area() int {
	return s.X * s.Y
}

func main() {
	v := Shape{3, 4}

	v.Area() // OK（3章と同様、Goが自動で &v に変換してくれる）

	var a Shaper

	a = &v // OK（*Shape は Shaper を満たす）
	//a = v  // NG！（Shape は Shaper を満たさない）

	fmt.Println(a.Area())
}
