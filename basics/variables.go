package basics

import "fmt"

func main() {
	var age int
	var name string = "Jhon"
	var name1 = "Jane"

	/* type inference in go with gopher simbol, this can only be used inside functions. Use inside
	functions and reasign values to variables with the same type
	*/
	count := 10
	lasName := "Smith"
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(count)
	fmt.Println(lasName)

	/*
		--- Default values:
		Numeric types: 0
		Boolean Types: false
		String Types: ""
		Pointers, Slices, Maps, Functions and Structs: nil

		--- Scope
		Variables in Go have block scope, for declare global variables use var keyword
	*/
}
