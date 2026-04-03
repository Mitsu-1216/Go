package main

import "fmt"

func main() {
	i, j := 324, 3232

	p := &i // アドレス取得
	fmt.Println("アドレス表示：", p)
	fmt.Println("値表示：", *p)
	*p = 48 // 値を代入
	fmt.Println("アドレス表示：", p)
	fmt.Println("値表示：", *p)
	fmt.Println("値表示：", i)

	p = &j // アドレス取得
	fmt.Println("値表示：", j)
	fmt.Println("値表示：", *p)
	*p = *p / 33 // 値を代入
	fmt.Println("アドレス表示：", p)
	fmt.Println("値表示：", *p)
	fmt.Println("値表示：", j)
}
