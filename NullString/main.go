package main

import (
	"database/sql"
	"fmt"
)

func main() {

	var name sql.NullString

	if name.Valid {
		fmt.Println(name.String) // DBの値が入っている
	} else {
		fmt.Println("NULL") // DBはNULL
	}

	var i int    // 0
	var s string // ""
	var b bool   // false

	var is []int         // nilスライス
	var m map[string]int // nil map

	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(is)
	fmt.Println(b)
	// m["a"] = 1 // これはpanic！

	if m == nil {
		m = make(map[string]int)
	}
	m["a"] = 1 // OK

	// nilスライスにそのまま追加できる
	is = append(is, 1) // OK（panicにならない）
}
