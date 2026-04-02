package main

import "fmt"

// func calc(sum int) (a, b int) {
// 	return
// }

func calc(sum1, sum2 int) (a, b int) {
	a = sum1 / sum2
	b = sum1 % sum2
	return
}

func main() {
	fmt.Println(calc(32, 24))
}
