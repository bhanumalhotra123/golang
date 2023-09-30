package main

import "fmt"

func main() {
	//slices
	characters := []string{"a", "b", "c", "d"} //looks like: ["a", "b", "c", "d"]
	//                                                 index:  0    1    2    3

	//loops
	for index, value := range characters {
		fmt.Println(index)
		fmt.Println(value)
		fmt.Println(characters[index])

	}

}

//range keyword with a slice (or array), it returns two values during each iteration:

//The index (position) of the current element.
//The value (element) at that index.
