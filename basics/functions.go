package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		1. publlic functions with uppercase letter
		2. private functions with lowercase letter
		3. if return is omitted functions will return zero values for their types
		4. arguments that are passed to a function are copied into the functions parameters,
		   modifications to that parameters inside a function do not affect original arguments
		5. anonimous functions are defined without a name
	*/
	greetAfunc := func() {
		fmt.Println("I am Anonymous function")
	}
	greetAfunc()
	/*
		we can use functions as types and functions can be assigned to variables, passed as
		arguments to other functions and returned from funcions, makem them a first class citizen
		or a first class object
	*/
	sum := add
	result := sum(18, 5)
	fmt.Println(result)

	// passing a function as an argument
	resultAplyOperation := applyOperation(2, 8, add)
	fmt.Println("result aply operation:", resultAplyOperation)

	// returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("result multiplyBy2:", multiplyBy2(6))

	//
	compareResult, err := compare(48, 789)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(compareResult)
}

func add(a, b int) int {
	return a + b
}

// function thattakes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "b is grater than b", nil
	} else {
		return "", errors.New("Unable to compare")
	}
}

/*
Go allows omitting the return values since they are already named in the function signature.
The function returns the current values of `quotient` and `remainder` without explicitly stating them.
*/
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// %v is a placeholder general for any value, %d for integers
