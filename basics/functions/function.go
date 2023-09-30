package main

import "fmt"

func main() {
	//functions
	sum := addTwoNumbers(1, 3)
	fmt.Println(sum)

}

func addTwoNumbers(left, right int) int {
	return left + right
}
