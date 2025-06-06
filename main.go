package main

import (
	"fmt"
)

func main() {
	score := 85
	if score > 90 {
		fmt.Println("Grade A")
	} else if score > 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C or Below")
	}

	if age := 20; age >= 18 {
		fmt.Println("You can consume alcohol")
	} else {
		fmt.Println("You can not consume alcohol")
	}

	if score := 80; score > 90 {
		fmt.Println("Grade A+")
	} else {
		fmt.Println("Grade B-")
	}

	lang := "go"

	switch lang {
	case "python":
		fmt.Println("Python dev")
	case "go":
		fmt.Println("go dev")
	default:
		fmt.Println("Unknown dev")
	}

	score = 90

	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 75:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("defer is done like a stack memory")

}
