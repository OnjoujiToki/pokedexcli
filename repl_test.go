package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World ",
			expected: []string{"hello", "world"},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected %v, got %v", cs.expected, actual)
		}
		for i, v := range actual {
			if v != cs.expected[i] {
				t.Errorf("Expected %v, got %v", cs.expected, actual)
			}
		}
	}
}
