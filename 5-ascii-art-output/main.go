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
	content,err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: banner file '%s' not found\n", fileName)
		os.Exit(1)
	}
	 
	// Split the file into a slice of individual lines
	rawString := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(rawString, "\n")

	bannerMap := make(map[rune][]string)
	character := ' '

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}
		bannerMap[character] = lines[i : i+8]
		character++
	}

	return bannerMap
}

func Render(input string, bannerMap map[rune][]string) string {
	var result strings.Builder
	input = strings.ReplaceAll(input, "\\n", "\n")
	if input == "" {
		return ""
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		 if line == "" {
			result.WriteString("\n")
			continue
		 }

		 for row := 0; row < 8; row++ {
			for _, char := range line {
				art, ok := bannerMap[char]
				if !ok {
					continue
				}
				result.WriteString(art[row])
			}
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