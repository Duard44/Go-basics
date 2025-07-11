package main

import "fmt"

func main() {
	message := "i promisse you this, i always look out for you"

	// print index and unicode value of string
	// for i, v := range message {
	// 	fmt.Println(i, v)
	// }

	// print index and unicode value of string
	for i, v := range message {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
	/*
		range iterates over a copy of the data structure it iterates over, modifying index or
		value inside the loop does not affect the original data structure

		1. for arrays ans slices 'range' interates in order from the first to last
		2. for maps interates over key:value pairs in an unspecified order
		3. for channels iterates until channel is closed
	*/
}
