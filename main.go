package main

import "fmt"

func main() {
	names := []string{"Shubham", "Amit", "Ravi"}
	for _, values := range names {
		fmt.Println("Values ", values)
	}
	for ind := range names {
		fmt.Println("Indexs ", ind)
	}
	for i := range names {
		if i==0 {
			names[i] = "Shubham Baheti"
		} 
	}
	for _, values := range names {
		fmt.Println("Values ", values)
	}
}
