package main

import "fmt"

func main() {
	fruits := "apple"

	switch fruits { // for multiple cases: switch apple, lemon, banana{... }
	case "apple":
		fmt.Println("It's an", fruits)
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknow fruit!")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greather than 1")
		fallthrough // is like not using break in other languages
	case num == 2:
		fmt.Println("Number is", num)
	default:
		fmt.Println("Not 2")
	}

	chechType(10)
	chechType(3.14)
	chechType("string")
	chechType(true)
}

// switch statements can be used with type assertionsto swich on the type of
// interface value
func chechType(x interface{}) {
	switch x.(type) { // when use type is not permited fallthrough
	case int:
		fmt.Println("It's an Integer")
	case float64:
		fmt.Println("It's an float64")
	case string:
		fmt.Println("It's an string")
	default:
		fmt.Println("Unknow type")
	}
}
