package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Bulbasaur Pikachu Charmander ",
			expected: []string{"bulbasaur", "pikachu", "charmander"},
		},
		{
			input:    "somethingforyourmind",
			expected: []string{"somethingforyourmind"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if len(word) != len(expectedWord) {
				t.Errorf("Expected: %s\nGot: %s", expectedWord, word)
			}
		}
	}
}
