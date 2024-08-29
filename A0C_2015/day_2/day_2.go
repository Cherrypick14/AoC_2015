package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

// CalculatePresentWrapping calculates the wrapping paper needed for a given dimension string
func CalculatePresentWrapping(dimensions string) int {
	// Split the dimension string by 'x'
	dimParts := strings.Split(dimensions, "x")
	if len(dimParts) != 3 {
		fmt.Printf("Invalid dimension format: %s\n", dimensions)
		return 0
	}

	// Convert string parts to integers
	l, err1 := strconv.Atoi(dimParts[0])
	w, err2 := strconv.Atoi(dimParts[1])
	h, err3 := strconv.Atoi(dimParts[2])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Invalid numbers in dimension: %s\n", dimensions)
		return 0
	}

	// Calculate areas of the sides
	lw := l * w
	wh := w * h
	hl := h * l

	// Calculate the smallest side area
    	smallestSide := min(lw, wh, hl)

		// Calculate the surface area of the box
		surfaceArea := 2*lw + 2*wh + 2*hl

		// Calculate the total amount of wrapping paper needed
		paperNeeded := surfaceArea + smallestSide
	
		return paperNeeded
	}

func main() {
	// Check the no of arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the input file.")
		return
	}
	// Grab the filename as argument

	fileName := os.Args[1]

	fmt.Println(CalculatePresentWrapping(fileName))
}


