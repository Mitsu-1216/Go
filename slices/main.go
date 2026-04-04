package main

import "fmt"

func main() {
	primes := [10]int{1, 8, 1, 8, 3, 2, 4, 4, 4, 0}
	fmt.Println(primes[1:4])
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(primes)
	fmt.Println(s)

}
