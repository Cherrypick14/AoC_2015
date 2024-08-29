package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

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



func TestCalculatePresentWrapping(t *testing.T) {
	// Read the dimensions input from file
	lines, err := readInput("input.txt")
	if err != nil {
		t.Fatalf("Error reading input file: %v", err)
	}

	// Generate test cases from the input lines
	tests := make([]struct {
		name     string
		line     string
		expected int
	}, len(lines))

	for i, line := range lines {
		tests[i] = struct {
			name     string
			line     string
			expected int
		}{
			name:     "test line " + strconv.Itoa(i),
			line:     line,
			expected: CalculatePresentWrapping(line),
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := CalculatePresentWrapping(tt.line)
			if total != tt.expected {
				t.Errorf("For input %s, expected %d but got %d", tt.line, tt.expected, total)
			}
		})
	}
}