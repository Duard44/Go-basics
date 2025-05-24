package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	result := a + b
	fmt.Println(result)

	result = a % b
	fmt.Println(result)

	const b1 float32 = 22 / 3.0 // it's ok because: int / float = float, int / int = int
	fmt.Println(b1)

	// Overflow with signed int
	var maxInt int64 = 9264684684354354656
	fmt.Println(maxInt)
	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned int
	var uMaxInt uint64 = 92646846843543546566
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow value exceeds minimun value that can be stored
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
