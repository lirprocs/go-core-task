package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
)

func sliceExample(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	for _, v := range arr {
		if v%2 == 0 {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func addElements(arr []int, i int) []int {
	arr = append(arr, i)
	return arr
}

func copySlice(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	return copyArr
}

func removeElement(arr []int, ind int) ([]int, error) {
	if ind < len(arr) && ind >= 0 {
		newArr := append(arr[:ind], arr[ind+1:]...)
		return newArr, nil
	}
	return nil, errors.New("Выход за пределы среза")
}

func main() {
	originalSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.IntN(100))
	}
	fmt.Println("Оригинальный слайс:", originalSlice)

	newSlice := sliceExample(originalSlice)
	fmt.Println("Четные числа:", newSlice)

	newSlice2 := addElements(newSlice, 5)
	fmt.Println("Слайс с добавленным элементом:", newSlice2)

	copiedSlice := copySlice(newSlice2)
	fmt.Println("Копия слайса:", copiedSlice)

	newSlice2 = append(newSlice2, 8)
	fmt.Println("Измененный оригинальный слайс:", newSlice2)
	fmt.Println("Копия после изменений в оригинале:", copiedSlice)
	if !reflect.DeepEqual(newSlice2, copiedSlice) {
		fmt.Println("Изменения в оригинальном слайсе не влияют на его копию.")
	}

	index := rand.IntN(len(copiedSlice))
	removedSlice, err := removeElement(copiedSlice, index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", index, removedSlice)
	}
}
