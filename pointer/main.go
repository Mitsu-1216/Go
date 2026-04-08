package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i        // ポインタに代入
	fmt.Println(p) // ポインタを表示
	*p = 21        // pの値に設定
	fmt.Println(i) // iの値を表示

	p = &j         // ポインタに代入
	*p = *p * 3    //
	fmt.Println(j) // jを表示

	t := &Vertex{1, 2}
	t.X = 100
	fmt.Println(t)
}
