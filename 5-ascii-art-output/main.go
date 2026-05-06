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
	
	// This reads the folder in which we have the banner file
	fileName = "banners/" + fileName
	// Read the raw file content
	content,err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: banner file '%s' not found\n", fileName)
		os.Exit(1)
	}
	 

	// This ensures that if you are working on Windows or the banner file has mixed line endings, the code won't break while trying to find the 8 lines of art.
	// Split the file into a slice of individual lines
	rawString := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(rawString, "\n")

	bannerMap := make(map[rune][]string)
	character := ' ' // The standard banner file starts at the Space character (ASCII 32) 

	// The standard banner format: Each character is 8 lines tall with 1 empty line above it.
	// Total height per character block = 9 lines.
	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break // Prevent "index out of range" if the file ends unexpectedly
		}
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


func main() {
	// This validates we have at least an argument from the command line
	// This is to make sure that the argument pass to the command line is not empty
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run . --output=<fileName.txt> something standard")
		return
	}

	var fileName string
	var remainingArgs []string

	// This is to check for the argument from the command-line input that has the flag and for extraction of file name
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			fileName = arg[len("--output="):]
		} else {
			remainingArgs = append(remainingArgs, arg)
		}
	}

	//Validate we have the text argument to print
	if len(remainingArgs) < 1 {
		fmt.Println("Usage: go run . --output=<fileName.txt> something standard")
	}

	userInput := remainingArgs[0]

	bannerStyle := "standard"
	if len(remainingArgs) > 1 {
		bannerStyle = remainingArgs[1]
	}

	bannerFile := bannerStyle + ".txt"

	// Initialize the font map and generate the art
	bannerMap := loadBanner(bannerFile)
	output := Render(userInput, bannerMap)

	// File creation and writing
	if fileName != "" {
		// 0644 provides 'Read/Write' permissions for the owner, 
		// and 'Read' permissions for everyone else.
		err := os.WriteFile(fileName, []byte(output), 0644)
	
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return

		}

		fmt.Printf("Success: ASCII art written to %s\n", fileName)

	} else {
		// Default behavior (output to terminal)
		fmt.Print(output)
	}
	
}