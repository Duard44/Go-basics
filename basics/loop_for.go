package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5, 6}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value:%d\n", index, value) // use Printf to use placeholders
		fmt.Printf("Index: %v, Value:%v\n", index, value) // v general value d for numbers
	}

	for i := 1; i <= 10; i++ {
		if i == 8 {
			break
		}

		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	// inner loops for spaces before stars
	// 	for j := 1; j <= rows; j++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // move top next line
	// }

	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off!")
}
