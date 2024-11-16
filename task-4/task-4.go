package main

import (
	"fmt"
)

func intersection(slice1, slice2 []string) []string {
	myMap := make(map[string]bool)
	var result []string

	for _, v := range slice1 {
		myMap[v] = true
	}

	for _, v := range slice2 {
		if myMap[v] {
			result = append(result, v)
			delete(myMap, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	slice := intersection(slice1, slice2)
	fmt.Println(slice)
}
