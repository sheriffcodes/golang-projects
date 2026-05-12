package main

import(
	"fmt"
	"strings"
)


// Declaring & initialising a map
// A map is Go's built-in key→value store. Two things matter before you touch a map: you must know its key and value types at compile time, and you must initialise it with make or a map literal — a nil map panics on write.

func main() {
	// ###Example 1
	// Method 1: make - zero value, ready to grow
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 82

	// Method 2: map literal
	cities := map[string]string{
		"NG": "Lagos",
		"JP": "Tokyo",
		"FR": "Paris",

	}
	fmt.Println(scores, cities)
	// map[Alice:95 Bob:82] map[FR:Paris JP:Tokyo NG:Lagos]

	// ###Test 1
	// Write a map literal of type map[string]float64 that stores three currency exchange rates (e.g. "USD", "EUR", "GBP"), then print the length of the map using the built-in len function. Paste your code and I'll review it.

	currency := map[string]float64{
		// Correct map literal syntax: map[string]float64{...}
		"USD": 20.1,
		"EUR": 20.2,
		"GBP": 20.3,
	}
	fmt.Println(len(currency))
	// 3
	fmt.Printf("%#v\n", currency)
	// The %#v verb is handy when you want to see the type alongside the values.
	// prints full type: map[string]float64{"EUR":20.2, ...}
	// map[string]float64{"EUR":20.2, "GBP":20.3, "USD":20.1}

	// ###Example 2
	// A simple comma-ok example
	permissions := map[string]bool{
		"read": true,
		"write": false, // explicitly set to false
	}
	// WITHOUT comma-ok — can't tell false from missing
	v1 := permissions["write"]
	fmt.Println(v1) // false — but is it set, or just missing?

	// WITH comma-ok — definitive answer
	if val1, ok := permissions["write"]; ok {
		fmt.Println("Write permission is explicitly set, value:", val1)
	}

	if _, ok := permissions["delete"]; !ok {
		fmt.Println("delete permissions is not in the map at all")
	}

	
	fmt.Println("//////////////////////////////////////")

	// ###Test 2
	// Given map[string]bool{"darkMode": true, "betaUI": false}, use comma-ok to print whether "betaUI" is explicitly set, and whether "offlineMode" is set at all.
	theme := map[string]bool {
		"darkMode": true,
		"betaUI": false,
	}

	v2 := theme["betaUI"]
	fmt.Println(v2)

	if val2, ok := theme["betaUI"]; ok {
		fmt.Println("betaUI permission is explicitly set, value:", val2)
		//betaUI permission is explicitly set, value: false
	} else {
		fmt.Println("betaUI is not in the map")
	}

	if _, ok := theme["offlineMode"]; !ok {
		fmt.Println("offlineMode not in the map")
		// offlineMode not in the map at all
	}

	// ###Example 3
	// Delete any word that appears only once — means you'll need to think about when to delete. You can't delete while ranging 
	// The word count challenge is a great one because it combines three things you now know: map literals, writing to maps, and iteration.
	colours := []string{"red", "blue", "red", "green", "blue", "red"}
	
	 //
	 counts := make(map[string]int)
	 for _, colour := range colours {
		
	 }
	
}