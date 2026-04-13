package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s1 []int  // nil
	s2 := []int{} // 空スライス

	// Go のスライスを JSON に変換する。nil は null、空スライスは [] になる。
	b1, _ := json.Marshal(s1)
	b2, _ := json.Marshal(s2)

	fmt.Println(string(b1)) // null
	fmt.Println(string(b2)) // []
}
