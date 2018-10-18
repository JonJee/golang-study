package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Println("add : ", add(1, 2))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
