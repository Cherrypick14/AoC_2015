// TestCalculatePresentWrapping verifies the CalculatePresentWrapping function.
package main

import (
	"testing"
	"strings"
)
func TestCalculatePresentWrapping(t *testing.T) {
	// Read the dimensions input from the file
	lines, err := readInput("input.txt")
	if err != nil {
		t.Fatalf("Error reading input file: %v", err)
	}

	for _, line := range lines {
		// Trim any surrounding whitespace
		line = strings.TrimSpace(line)

		// Split the line into dimensions
		dimensions := line

		// Split dimensions and parse the expected values
		dimParts := strings.Split(dimensions, "x")
		if len(dimParts) != 3 {
			t.Errorf("Invalid dimension format: %s", dimensions)
			continue
		}

		// For demonstration, using a map of expected results as in the previous example
		expectedResults := map[string]struct {
			paperNeeded  int
			ribbonNeeded int
		}{
			"5x21x29": {5088, 178},
			"9x21x16": {2082, 432},
			"9x6x10":  {552, 198},
			"9x6x4":   {288, 108},
			"24x14x29":{2488, 1948},
			"28x11x6": {1692, 504},
			"10x22x1": {484, 32},
			"21x30x20":{3180, 1180},
			"13x17x8": {1782, 520},
			// Add more known results here...
		}

		expected, found := expectedResults[dimensions]
		if !found {
			t.Errorf("No expected results for dimensions: %s", dimensions)
			continue
		}

		paperNeeded, ribbonNeeded := CalculatePresentWrapping(dimensions)
		if paperNeeded != expected.paperNeeded {
			t.Errorf("For dimensions %s, expected paper %d but got %d", dimensions, expected.paperNeeded, paperNeeded)
		}
		if ribbonNeeded != expected.ribbonNeeded {
			t.Errorf("For dimensions %s, expected ribbon %d but got %d", dimensions, expected.ribbonNeeded, ribbonNeeded)
		}
	}
}