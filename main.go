package main

import "fmt"

// Make a Rectangle struct with Width and Height. Add a method:

// Area() float64 → returns area

// Scale(f float64) → scales both width and height

type Rectangle struct {
	Height, Width float64
}

func (rec *Rectangle) Area() float64 {
	area := rec.Height * rec.Width
	return area
}

func (rec *Rectangle) Scale(f float64) {
	rec.Height = rec.Height * f
	rec.Width = rec.Width * f
}

func main() {
	rec := Rectangle{3, 4}
	fmt.Printf("the initial height=%.2f and width=%.2f\n", rec.Height, rec.Width)
	rec.Scale(5)
	fmt.Printf("after scale height=%.2f and width=%.2f\n", rec.Height, rec.Width)
	fmt.Println("The area is ", rec.Area())

}
