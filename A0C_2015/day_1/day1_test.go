package main

import "testing"


func TestSanta(t *testing.T) {

	tests := []struct {
		input  string
		expectedCount int
		expectedpos int
	} {
		{"(())", 0, -1},
        {"()())", -1, 5},
        {")", -1, 1},
	}

	for _, test := range tests {
		  count, pos := Santa(test.input)
		   if count != test.expectedCount || pos != test.expectedpos {
			 t.Errorf("For input %q, expected (%d, %d), got (%d, %d)", test.input, test.expectedCount, test.expectedpos,count,pos)
		   }
	}
}