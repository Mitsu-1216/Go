package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0で割ることができません")
	}
	return a / b, nil
}
 
func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
