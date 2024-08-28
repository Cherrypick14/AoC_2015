package main

import (
	"fmt"
	"os"
)

func Santa(stars string) (int, int) {

	count := 0

	firstBasementPos := -1

	for i, char := range stars {
		if char == '(' {
			count++

		} else if char == ')' {
			count--

		}

		//check if santa is at the basement (-1) && the firstbasement occurrence is 1

		if count == -1 && firstBasementPos == -1 {
			firstBasementPos = i + 1 // Convert to 1-based index
		}
	}

	return count, firstBasementPos
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Out of bounds")
	} else if len(os.Args) == 2 {

		args1 := os.Args[1]

		fmt.Println(Santa(args1))
	}

}
