package main

import "fmt"

func main() {
	// conditional statements
	passingGrade := 65
	grade := 72

	if grade >= passingGrade {
		fmt.Println("You passed!")
	} else {
		fmt.Println("You failed!")
	}
}
