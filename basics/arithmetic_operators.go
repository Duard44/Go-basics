package main

import "fmt"

func main() {
	var a, b int = 10, 3
	result := a + b
	fmt.Println(result)

	result = a % b
	fmt.Println(result)

	const b1 float32 = 22 / 3.0 // it's ok because -> int / float = float
	fmt.Println(b1)
}
