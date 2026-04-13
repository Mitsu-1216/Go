package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	// %w でErrNotFoundをラップして返す
	// → 「findUser:」という文脈情報を付けつつ、元のエラーも保持できる
	return fmt.Errorf("findUser: %w", ErrNotFound)
}
func main() {

	err := findUser(1)

	fmt.Println(err == ErrNotFound)          // false（ラップされているので一致しない）
	fmt.Println(errors.Is(err, ErrNotFound)) // true（中身まで確認してくれる）

}
