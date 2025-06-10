package main

import "fmt"

func indexOfNum(s []int, x int) int {
	for index, val := range s {
		if val == x {
			return index
		}
	}
	return -1
}

func indexOfStr(s []string, temp string) int {
	for index, val := range s {
		if val == temp {
			return index
		}
	}
	return -1
}

func indexOfAll[T comparable](slice []T, value T) int{
	for index,val :=range slice{
		if val == value {
			return index
		}
	}
	return -1;
}

func main() {
	s := []int{3, 4, 5, 6, 7}
	fmt.Println(indexOfAll(s, 7))
	temp := []string{"Hello", "World", "Mohan", "Ram"}
	fmt.Println(indexOfAll(temp, "Mohan"))
}
