package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "blitz",
			expected: []string{"blitz"},
		},
		{
			input:    "this is a test of many words",
			expected: []string{"this", "is", "a", "test", "of", "many", "words"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of return slice (%d) doesn't match expected length (%d)", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s doesn't match %s", word, expectedWord)
			}
		}
	}
}
