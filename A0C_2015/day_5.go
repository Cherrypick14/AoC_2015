package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

// isNice checks if a string is "nice" based on new criteria.
func isNice(str string) int {
	if !hasNonOverlappingPair(str) {
		return 0
	}
	if !hasRepeatedLetterWithOneBetween(str) {
		return 0
	}
	return 1
}

// hasNonOverlappingPair checks for a pair of letters appearing twice without overlapping.
func hasNonOverlappingPair(str string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		pair := str[i:i+2]
		if pos, exists := pairs[pair]; exists && i > pos+1 {
			return true
		}
		pairs[pair] = i
	}
	return false
}

// hasRepeatedLetterWithOneBetween checks for a letter repeating with exactly one letter in between.
func hasRepeatedLetterWithOneBetween(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalNiceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) == 1 {
			totalNiceCount++
		}
	}

	fmt.Println("The total number of nice strings is", totalNiceCount)
}
