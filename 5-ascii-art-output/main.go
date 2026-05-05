package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(fileName string) map[rune][]string {
	content,err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: banner file '%s' not found\n", fileName)
		os.Exit(1)
	}

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
			fileName = arg[len("--output"):]
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

	bannerFile := bannerStyle + "txt"

	// Initialize the font map and generate the art
	bannerMap := loadBanner(bannerFile)
	output := Render(userInput, bannerMap)

	// File creation and writing
	if fileName != "" {
		err := os.WriteFile(fileName, []byte(output), 0644)
	
		if err != nil {
			fmt.Println("Error writing to file: %v\n", err)
			return

		}

		fmt.Println("Success: ASCII art written to %s\n", fileName)

	} else {
		// Default behavior (output to terminal)
		fmt.Print(output)
	}
	
}