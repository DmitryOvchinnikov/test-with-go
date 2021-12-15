package main

import (
	"fmt"

	"github.com/dmitryovchinnikov/test-with-go/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: Wanted 11, but got %d", sum)
		panic(msg)
	}
	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: Wanted 15, but got %d", add)
		panic(msg)
	}
	fmt.Println("PASS")
}
