package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
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

// CalculatePresentWrapping calculates the wrapping paper and ribbon needed for a given dimension string
func CalculatePresentWrapping(dimensions string) (int, int) {

	// Trim any extra spaces from the dimension string
	dimensions = strings.TrimSpace(dimensions)
	
	// Split the dimension string by 'x'
	dimParts := strings.Split(dimensions, "x")
	if len(dimParts) != 3 {
		fmt.Printf("Invalid dimension format: %s\n", dimensions)
		return 0, 0
	}

	// Convert string parts to integers
	l, err1 := strconv.Atoi(dimParts[0])
	w, err2 := strconv.Atoi(dimParts[1])
	h, err3 := strconv.Atoi(dimParts[2])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Invalid numbers in dimension: %s\n", dimensions)
		return 0, 0
	}

	// Calculate areas of the sides
	lw := l * w
	wh := w * h
	hl := h * l

	// Calculate the smallest side area
	smallestSide := min(lw, wh, hl)

	// Calculate the surface area of the box
	surfaceArea := 2*lw + 2*wh + 2*hl

	// Calculate the smallest perimeter of any face
	perimeter := 2 * min(l+w, w+h, h+l)

	// Calculate the volume of the box
	volume := l * w * h

	// Calculate the total amount of wrapping paper needed
	paperNeeded := surfaceArea + smallestSide

	// Calculate the total ribbon required (perimeter for wrapping + volume for bow)
	totalRibbonRequired := perimeter + volume

	return paperNeeded, totalRibbonRequired
}

// Read the input file and return a slice of strings
func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	// Check the number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the input file.")
		return
	}
	// Grab the filename as an argument
	fileName := os.Args[1]

	// Read the dimensions from the file
	lines, err := readInput(fileName)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}
     
	totalRibbon := 0
	// Process each line from the file
	for _, line := range lines {
		paperNeeded, ribbonNeeded := CalculatePresentWrapping(line)
		totalRibbon += ribbonNeeded
		fmt.Printf("Dimensions: %s\nPaper Needed: %d\nRibbon Needed: %d\n", line, paperNeeded, ribbonNeeded)
		
	}
	// Print the total ribbon needed
	fmt.Printf("Total Ribbon Needed: %d\n", totalRibbon)
}

