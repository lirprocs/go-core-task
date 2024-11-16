package main

import (
	"fmt"
)

func intersection(slice1, slice2 []int) ([]int, bool) {
	myMap := make(map[int]bool)
	var result []int
	//resultBool := false

	for _, v := range slice1 {
		myMap[v] = true
	}

	for _, v := range slice2 {
		if myMap[v] {
			result = append(result, v)
			delete(myMap, v)
		}
	}

	if len(result) > 0 {
		return result, true
	} else {
		return result, false
	}

}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	slice, result := intersection(a, b)
	fmt.Println(result, slice)
}
