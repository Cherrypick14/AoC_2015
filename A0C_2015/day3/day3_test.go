package main

import (
	"testing"
)

func TestDelivery(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "^v", expected: 3},
		{input: "^>v<", expected: 3},
		{input: "^v^v^v^v^v ", expected: 11},
	}

	for _, test := range tests {
		count := Delivery(test.input)
		if count != test.expected {
			t.Errorf("Wrong output: got %d expected %d", count, test.expected)
		}
	}
}
