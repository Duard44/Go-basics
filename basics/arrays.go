package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 6
	fmt.Println(numbers)

	fruits := [4]string{"apple", "banana", "orange", "lemon"}
	fmt.Println(fruits)
	fmt.Println("Second element:", fruits[1])

	/*
		When you assign an array to a new variable or pass an array as an argument to a function,
		a copy of the original array is created and modifications to the copy do not affect the
		original array.
	*/

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 3

	fmt.Println("original array:", originalArray)
	fmt.Println("copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, ":", numbers[i])
	}

	// use blank indentifier _ to avoid use variable
	for i, v := range numbers { // index, value, _ ingnore
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	// blank indentifier store unused values
	x := 4
	_ = x

	// compare arrays
	arry1 := [3]int{5, 6, 7}
	arry2 := [3]int{5, 6, 7}

	fmt.Println("array 1 is uqual to array 2??", arry1 == arry2)

	// multidimensional arrays
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
	copyngArray()

}

func copyngArray() {
	/*
		by default, when you assign the variable to another variable,
		a copy of that variable is assigned to the new variable.
		if we were to use the original array in copied array, we
		would have to use pointers and addresses.
		addresses are denoted by ampersand sign and pointers are denoted
		by the asterisk sign
	*/
	originalArray := [3]int{1, 2, 3}

	/*
		now copied array is not a variable which stores the duplicate values of original array.
		is actually original array because copied array is a pointer,
		and this pointer now holds the address of the original array
	*/
	var copiedArray *[3]int
	copiedArray = &originalArray

	copiedArray[1] = 400

	fmt.Println("original array:", originalArray)
	fmt.Println("copied array:", copiedArray)
}
