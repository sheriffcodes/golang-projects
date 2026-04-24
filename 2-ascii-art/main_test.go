package main

import (
	"testing"
)

func TestRender(t *testing.T) {
	// Need a mock bannerMap for testing so as to not rely on standard.txt
	mockBanner := map[rune][]string{
		'A': {
			" AAA ",
			"A   A",
			"AAAAA",
			"A   A",
			"A   A",
			"     ",
			"     ",
			"     ",
		},
		'B': {
			"BBBB ",
			"B   B",
			"BBBB ",
			"B   B",
			"BBBB ",
			"     ",
			"     ",
			"     ",
		},
	}

	// Defining our test cases (The Table)
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "Single Character",
			input: "A",
			expected: " AAA \nA   A\nAAAAA\nA   A\nA   A\n     \n     \n     \n",
		},
		{
			name:  "Horizontal Concatenation",
			input: "AB",
			// Row 1 of A + Row 1 of B, then Row 2...
			expected: " AAA BBBB \nA   AB   B\nAAAAABBBB \nA   AB   B\nA   ABBBB \n          \n          \n          \n",
		},
		{
			name:  "Handling Newlines",
			input: "A\\nB",
			expected: " AAA \nA   A\nAAAAA\nA   A\nA   A\n     \n     \n     \nBBBB \nB   B\nBBBB \nB   B\nBBBB \n     \n     \n     \n",
		},
		{
			name:  "Empty Input",
			input: "",
			expected: "",
		},
	}

	// Iterating through the table
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Render(tc.input, mockBanner)
			if got != tc.expected {
				t.Errorf("Render(%q) failed.\nGot:\n%q\nExpected:\n%q", tc.input, got, tc.expected)
			}
		})
	}
}