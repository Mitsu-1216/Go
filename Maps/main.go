package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	m := make(map[string]int)

	for _, w := range words {
		m[w]++
	}

	return m
}
func main() {
	wc.Test(WordCount)
	fmt.Println(WordCount("I am I am Go"))
}

type error interface {
	Error() string
}
