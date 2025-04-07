package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "                 ",
			expected: []string{},
		},
		{
			input:    "          j       ",
			expected: []string{"j"},
		},
		{
			input:    "a                 b",
			expected: []string{"a", "b"},
		},
		{
			input:    "Hello    456     !!        b",
			expected: []string{"hello", "456", "!!", "b"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) != len(c.expected), %v != %v", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word does not match, expected: %s, got: %s", actual[i], c.expected[i])
				t.Fail()
			}
		}
	}
}
