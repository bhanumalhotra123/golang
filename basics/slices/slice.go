package main

import "fmt"

func main() {
	//slices
	characters := []string{"a", "b", "c", "d"} //looks like: ["a", "b", "c", "d"]
	//      index:  0    1    2    3
	firstChar := characters[0]
	lastChar := characters[3]

	fmt.Println(firstChar)
	fmt.Println(lastChar)

}
