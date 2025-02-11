package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello Fine hair",
			expected: []string{"hello", "fine", "hair"},
		},
	}

	for _, c := range cases {
		output := cleanInput(c.input)

		if len(output) != len(c.expected) {
			t.Errorf("incorrect output")
		}

		for i := range output {
			if output[i] != c.expected[i] {
				t.Errorf("incorrect output")
			}
		}
	}
}