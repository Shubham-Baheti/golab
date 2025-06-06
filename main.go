package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "shubham"
	a[1] = "bhanu"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 56, 7, 8}
	fmt.Println(primes)
}
