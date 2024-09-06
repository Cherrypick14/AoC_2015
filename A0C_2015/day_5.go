package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Santa needs help figuring out which strings in his text file are naughty or nice.

// Criteria for nice strings:
// > Contains at least 3 vowels (aeiou)
// > Contains at least one letter that appears twice in a row (xx, abcdde, aabbccdd)
// > Should not contain the strings (ab, cd, pq, xy)

// isNice checks if a given string meets the criteria for being "nice".
func isNice(str string) int {
	// List of forbidden substrings
	nonstrings := []string{"ab", "cd", "pq", "xy"}

	// Check if the string contains any of the forbidden substrings
	for _, forbidden := range nonstrings {
		if strings.Contains(str, forbidden) {
			return 0 // Return 0 if the string is not nice
		}
	}

	vowelCount := 0 // To count vowels in the string
	hasDouble := false // To check if there's a letter that appears twice in a row

	// Iterate over the string to count vowels and check for double letters
	for i := 0; i < len(str); i++ {
		if isVowel(rune(str[i])) {
			vowelCount++	
		}

		if i > 0 && str[i] == str[i-1] {
			hasDouble = true
		}
	}
   
	// Check if the string meets the criteria for being nice
	if vowelCount >= 3 && hasDouble {
		return 1 // Return 1 if the string is nice
	}
	return 0 // Return 0 if the string is not nice
}

// isVowel checks if a given rune is a vowel.
func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func main() {
	// Open the file named "data.txt"
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when main function exits

	// Create a buffered reader and scanner for the file
	filename := bufio.NewReader(file)
	newfile := bufio.NewScanner(filename)

	var output int // To store the result of isNice function
	totalNiceCount := 0 // To count the total number of nice strings

	// Read each line from the file and check if it's nice
	for newfile.Scan() {
		line := newfile.Text()
		output = isNice(line)
		totalNiceCount += output
	}

	// Print the total number of nice strings
	fmt.Println("The total number of nice strings is", totalNiceCount)
}
