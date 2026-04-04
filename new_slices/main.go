package main

import "fmt"

func main() {
	// 初期スライス（sはcapを小さくする）
	s := make([]int, 3, 3)
	b := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3
	b[0], b[1], b[2] = 1, 2, 3

	// コピー
	t := s

	fmt.Println("=== append前 ===")
	fmt.Printf("s: %v, len=%d, cap=%d, addr=%p\n", s, len(s), cap(s), &s[0])
	fmt.Printf("b: %v, len=%d, cap=%d, addr=%p\n", b, len(b), cap(b), &b[0])
	fmt.Printf("t: %v, len=%d, cap=%d, addr=%p\n", t, len(t), cap(t), &t[0])

	// append（sの容量をオーバーさせる）
	s = append(s, 4)
	b = append(b, 4)

	fmt.Println("\n=== append後 ===")
	fmt.Printf("s: %v, len=%d, cap=%d, addr=%p\n", s, len(s), cap(s), &s[0])
	fmt.Printf("b: %v, len=%d, cap=%d, addr=%p\n", b, len(b), cap(b), &b[0])
	fmt.Printf("t: %v, len=%d, cap=%d, addr=%p\n", t, len(t), cap(t), &t[0])
}
