package main

import (
	"fmt"
	"math/rand"
)

func swap(a, b string) (string, string) {
	return b, a
}

func valueAssign(sum int) (int, int) {
	var a int = sum * 4 / 9
	var b int = sum*5 + 4
	return a, b
}

func main() {
	fmt.Println("Hello from goLearning")
	fmt.Println("This is a random number:", rand.Intn(100))
	a, b := swap("hello", "world")
	println("Swapped values:", a, b)
	x, y := valueAssign(17)
	fmt.Println("Assigned values:", x, y)

}
