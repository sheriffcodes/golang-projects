
package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestOutputFile(t *testing.T) {
	// 1. Define our test parameters
	outputFile := "test_result.txt"
	inputText := "Hello"
	banner := "standard"

	// Cleanup: Remove the test file if it exists from a previous run
	os.Remove(outputFile)

	// 2. Simulate: go run . --output=test_result.txt "Hello" standard
	cmd := exec.Command("go", "run", ".", "--output="+outputFile, inputText, banner)
	
	// Capture any errors from running the command
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Program failed to run: %v", err)
	}

	// 3. Verify the file was created
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Expected file %s was not created", outputFile)
	}

	// 4. Basic content verification
	// We know "Hello" in standard ASCII art should have 8 lines (plus potentially a trailing newline)
	lines := strings.Split(string(content), "\n")
	
	// 'Hello' shouldn't be empty, and it should have multiple lines of ASCII art
	if len(content) == 0 {
		t.Error("The output file is empty, but should contain ASCII art")
	}

	// Standard banner art is 8 lines tall per text line
	if len(lines) < 8 {
		t.Errorf("Expected at least 8 lines of ASCII art, got %d", len(lines))
	}

	// Cleanup after test
	os.Remove(outputFile)
}