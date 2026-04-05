package main

import "fmt"

// インタフェース
type Calc interface {
	Add(a int) int
}

// 構造体
type User struct {
	name string
	age  int
}

// メソッド（インタフェースを満たす）
func (u User) Add(a int) int {
	return u.age + a
}

func main() {
	var c Calc = User{name: "太郎", age: 20}
	fmt.Printf("来年は%d歳です\n", c.Add(1))
}
