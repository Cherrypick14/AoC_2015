package main

import (
	"testing"
)

func TestFindAdventCoin(t *testing.T){
	tests := []struct {
		input string
		expected int
	}{
       {input: "bgvyzdsv", expected:1038736},
	}

	for _, test := range tests {
		count := FindAdventCoin(test.input)
		 if count != test.expected {
			 t.Errorf("Wrong count %d",test.expected)
		 }
	}
}