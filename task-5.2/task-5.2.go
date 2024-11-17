package main

import (
	"fmt"
	"math/rand"
)

func random(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go random(ch, 10)

	for val := range ch {
		fmt.Println(val)
	}
}
