package main

func calc(a, b int) (int, int) {
	return a * 10, b / 10
}

func main() {
	println(calc(100, 200))
}
