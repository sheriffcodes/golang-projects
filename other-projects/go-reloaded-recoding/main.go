package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	// Q1
	// Write a function that capitalizes the first letter of each word
	// "hello, ade!"
	inputq1 := "hello! ade"
	fmt.Println(ToCap(inputq1))

	// Q2
	// Write a function that splits a string into words
	// "go is fun" --> ["go", "is", "fun"]
	inputq2 := "hello! ade"
	fmt.Println(ToSplit(inputq2))

	// Q3
	// Write a function that joins words into a sentence
	// ["go", "is", "fun"] --> "go is fun"
	inputq3 := []string{"go", "is", "fun"}
	fmt.Println(ToJoin(inputq3))

	// Q4
	// Write a function that processes a full sentence and fixes/changes all "a" -> "an"
	// "An amazing boy. A honest gal!"
	inputq4 := "An amazing boy. A honest gal!"
	fmt.Println(ChangeToAn(inputq4))

	// Q5
	// Write a function that fixes spaceing inside single quotes
	// "'  Hello World  '" --> "'Hello World'"
	inputq5 := "'  Hello World  '"
	fmt.Println(FixSingleQuote(inputq5))

	// Q6
	// Explain what strings.Fields() does and why it's useful in this project

	// Q7
	// Why do we use strconv.ParseInt(s, 16, 64) instead of writing our own hex converter?

	// Q8
	// Write a function that removes leading and trailing spaces without using strings.TrimSpace
	// " hello " -> "  hello"
	inputq8 := "  Hi Shade  "
	fmt.Println(TrimSpace(inputq8))

	// Q9
	// Write a function that shifts each letter in a string by n positions in the alphabet (wrapping around). Non-letters are unchanged.
	inputq9 := "xyza"
	fmt.Println(NextThreePos(inputq9, 3))

	// Q10
	// Write a function that removes vowels from a string
	// "Hello World" -> "Hll Wrld"
	inputq10 := "Hello World"
	fmt.Println(RemoveVow(inputq10))
}


// A1
func ToCap(txt string) string {
	words := strings.Fields(txt)

	for i, word := range words {

		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

// A2
func ToSplit(txt string) []string {
	words := strings.Fields(txt)
	return words
}

// A3
func ToJoin(txt []string) string {
	words := strings.Join(txt, " ")
	return words
}

// A4
func ChangeToAn(txt string) string {
	words := strings.ReplaceAll(txt, "A", "An")
	return words
}

// A5
func FixSingleQuote(txt string) string {
	re := regexp.MustCompile(`['"]\s*(.*?)\s*["']`)
	return re.ReplaceAllString(txt, "'$1$2$3'")
}

// A6
// strings.Fields() is a standard go package that splits a string into a slice of strings using whitespace as the separator. it handles: spaces " " | tabs "\t" | newlines "\n" 
//
// words := strings.Fields("go is fun") ----> []string{"go", "is", "fun"}

// A7
// strconv.ParseInt is a standard go package that converts string → number | understands bases (like hex, binary, decimal) | returns an integer + error ---> num, err := strconv.ParseInt("1A", 16, 64)
// num = 26

// A8
func TrimSpace(txt string) string {
	words := strings.Fields(txt)
	return strings.Join(words, " ")
}

// A9
// In Go, when you use range on a string: ch in {for _, ch := range txt} is of type rune | A rune is just an alias for int32 | It represents a Unicode character
//Why 26? Because the alphabet has 26 letters. shifting by 27 → same as shifting by 1. So: 29 % 26 = 3
// 'h' in "hello" = 7 (next 7 letters from 'a') | 8 - 1(a) = 7 (“h is 7 letters after a, so 'a' is now 0”)| 7 + 3(pos) + 26(to positive) = 36 | 36 % 26 = 10 | 10 + 'a' (1) = 11 | letter at position 11 = 'k' or Next 10 after 'a' = 'k' (adding 10 + 'a' = 'k' to convert back to letter)
func NextThreePos(txt string, pos int) string {
	result := ""
	pos = pos % 26

	for _, ch := range txt {

		if ch >= 'a' && ch <= 'z' {
			shifted := ((ch-'a')+rune(pos)+26)%26 + 'a'
			result += string(shifted) // else if for UpperCase
		} else {
			result += string(ch)
		}
	}

	return result
}

// A10
func RemoveVow(txt string) string {
	result := ""

	for _, char := range txt {
		if char
	}
}