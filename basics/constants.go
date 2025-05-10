package basics

import "fmt"

const PI = 3.14 //untype constant
const GRAVITY = 9.81

func main() {
	const days int = 7
	// use cont block to make your code more readable
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4 // type constant
	)
	fmt.Println(PI)
	fmt.Println(GRAVITY)
}
