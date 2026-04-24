package main

import (
	"fmt"
	"os"
	"strings"
)

// THE STORAGE ENGINE (LOADBANNER)
// loadBanner converts the banner file into a searchable dictionary.
// It maps each character (rune) to its 8-line ASCII representation.
func loadBanner(fileName string) map[rune][]string {

	// Read the raw file content
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading banner file:", err)
		os.Exit(1)
	}

	// Split the file into a slice of individual lines
	lines := strings.Split(string(content), "\n")
	
	bannerMap := make(map[rune][]string)
	character := ' ' // The standard banner file starts at the Space character (ASCII 32)

	// Populate the map
	// We jump 9 lines at a time: 1 empty separator line + 8 lines of art
	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}
		// Store exactly the 8 lines of ASCII art for the current character
		bannerMap[character] = lines[i : i+8]
		character++
	}

	return bannerMap
}

// THE PROCESSING ENGINE (RENDER)
// Render processes the input string and builds the final multi-line output string.
func Render(input string, bannerMap map[rune][]string) string {
	var result strings.Builder

	// Replace literal backslash+n ("\\n") with a real newline character ("\n")
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Special case: empty input
	if input == "" {
		return ""
	}

	// Split the input into segments whenever a newline is encountered
	lines := strings.Split(input, "\n")

	for _, line := range lines {

		// Avoid adding an unnecessary blank line if the input ends with a newline
		// If a segment is empty (from a double newline like "A\n\nB"), add a blank line
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// The Core Printing Logic:
		// Because we print to the terminal line-by-line (horizontally), 
		// we must print the first row of every character in the word, 
		// then the second row, and so on.
		for row := 0; row < 8; row++ {
			for _, char := range line {
				// Safety check: verify the character exists in our loaded font
				art, ok := bannerMap[char]
				if !ok {
					continue // Ignore characters not supported by the banner file
				}
				// Append the specific row of this character to our builder
				result.WriteString(art[row])
			}
			// Once all characters have finished printing this specific row, add a newline
			result.WriteString("\n")
		}
	}

	return result.String()
}

// THE EXECUTION ENGINE (MAIN)
func main() {
	// Ensure the user provided a string to convert
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [string]")
		return
	}

	// Capture user input from command line arguments
	input := os.Args[1]

	// Initialize the font map and generate the art
	bannerMap := loadBanner("standard.txt")
	output := Render(input, bannerMap)

	// Final output to the terminal
	fmt.Print(output)
}