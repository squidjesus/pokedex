package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO w0rLD",
			expected: []string{"hello", "w0rld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("actual: %v", actual)
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) %d != len(expected) %d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s != %s", word, expectedWord)
				continue
			}
		}

	}
}
