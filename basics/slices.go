package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)

	/*
		slices are references to underlying arrays.

		They do not store any data themselves but provide a window into the array's elements.

		Slices can grow and shrink dynamically.

		And we have the same function Len which can check the length of the slice.
	*/

	a := [5]int{1, 2, 3, 4, 5}

	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("sliceCopy", sliceCopy)

	// nill slices does not reference any underlying array
	var nilSlice []int
	_ = nilSlice

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println(slice1[3])
	slice1[3] = 66
	fmt.Println(slice1[3])

	// compare slices
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	// multidimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
