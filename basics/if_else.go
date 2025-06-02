package main

import "fmt"

func main() {
	score := 80

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
		fmt.Println("...")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
		fmt.Println("...")
	}

	score2 := 90

	if score2 >= 90 {
		fmt.Println("Grade A")
	}

	if score2 >= 80 {
		fmt.Println("Grade B")
	}

	if score2 >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
}
