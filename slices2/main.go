package main

import "fmt"

func main() {
	arrays := [10]int{1, 8, 1, 8, 3, 2, 4, 4, 4, 0} //配列
	fmt.Println(arrays)

	slices := []bool{true, false} //スライス
	fmt.Println(slices)
}
