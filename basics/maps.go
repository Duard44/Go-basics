package main

import (
	"fmt"
	"maps"
)

func main() {
	/*
		maps store and retrieve key:value pairs, each key will be unique withon the map
		maps are unorderer collections of key value pairs meaning thats not garanteed order
		when iterating over them
	*/
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 16
	myMap["key2"] = 17

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key2"])
	fmt.Println(myMap["key3"]) // id the key doesn't exist zero value int and "" if is string
	myMap["key2"] = 18
	fmt.Println(myMap["key2"])

	delete(myMap, "key1")
	fmt.Println(myMap["key1"])

	// the second value is optional and is named 'ok' for convention
	value, ok := myMap["key2"]
	fmt.Println("value:", value, " unknowvalue:", ok)
	fmt.Println("Is a value asosociated with key2?", ok)
	clear(myMap)
	fmt.Println(myMap)

	fmt.Println("-----------------------------------------------")
	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7}
	fmt.Println(myMap2)

	fmt.Println("maps are equal?", maps.Equal(myMap, myMap2))

	fmt.Println("-----------------------------------------------")
	fmt.Println("Iterate over maps")

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}

	// a map that hasn't been initialized but only declared it is initialized to a 'nil' value
	var myMap3 map[string]string
	fmt.Println(myMap3 == nil)

	myMap3 = make(map[string]string)

	stringValueMyMap3 := myMap3["value1"]
	fmt.Println(stringValueMyMap3)

	myMap3["value1"] = "10"
	fmt.Println(myMap3["value1"])

	// lenght of map
	fmt.Println("-----------------------------------------------")
	fmt.Println("length of a map")
	fmt.Println("myMap3 length is:", len(myMap3))

	fmt.Println("-----------------------------------------------")
	fmt.Printf("nested maps")

	nestedMap := make(map[string]map[string]string)

	nestedMap["map1"] = myMap3
	fmt.Println(nestedMap)
}
