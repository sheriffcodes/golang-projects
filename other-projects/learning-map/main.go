package main

import(
	"fmt"
	"sort"
)

// A Map is one of the most powerful tools in coding. Think of it like a dictionary or a lookup table.
// scores := make(map[string]int)
// scores := make(map[key]value)
// The Key is the name (the string).
// The Value is the score (the int, or integer).
// Declaring & initialising a map
// A map is Go's built-in key→value store. Two things matter before you touch a map: you must know its key and value types at compile time, and you must initialise it with make or a map literal — a nil map panics on write.

func main() {
	// ###Example 1
	// Method 1: make will give it a zero value (ready to grow), rather than nil
	// I explained more about "make()" in Test 4
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
	// Write a map literal of type map[string]float64 that stores three currency exchange rates (e.g. "USD", "EUR", "GBP"), then print the length of the map using the built-in len function.

	currency := map[string]float64{
		// Correct map literal syntax: map[string]float64{...}
		"USD": 20.1,
		"EUR": 20.2,
		"GBP": 20.3,
	}
	fmt.Println("Length of map = %v", len(currency))
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
	// Write permission is explicitly set, value: false
	// Delete permissions is not in the map at all
	
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
	fmt.Println("//////////////////////////////")
	// Delete any word that appears only once — means you'll need to think about when to delete. You can't delete while ranging 
	// The word count challenge is a great one because it combines three things you now know: map literals, writing to maps, and iteration.
	colours := []string{"red", "blue", "red", "green", "blue", "red"}
	
	 //
	 counts := make(map[string]int)
	 for _, colour := range colours {
		counts[colour]++ // zero value of int is 0, so this is safe on first hit
		// This is a Loop. It goes through your list of colors one by one.
		// #The Magic Part: In Go, if you ask a map for a key that isn't there yet, it gives you a "Zero Value." For numbers, that's 0.
		// #So, the first time the code sees "red," it says "Red's count is 0, let's add 1." The next time it sees "red," it says "Red is already at 1, let's make it 2."
	 }
	 fmt.Println("before delete:", counts)
	 // map[blue:2 green:1 red:3]

	 // Deleting entries that appear only once
	 for key, value := range counts {
		if value == 1 {
			delete(counts, key)
		}
	 }
	fmt.Println("after delete:", counts)
	// map[blue:2 red:3]
	// Three things to notice:

	// #counts[colour]++ works even on a key that doesn't exist yet — Go returns the zero value (0) for a missing int key, then increments it to 1. No need to check first.
	// #You can safely delete inside a range loop in Go — the spec explicitly allows it. The deleted key simply won't appear in subsequent iterations.
	// #The two-pass approach (count first, then delete) is still cleaner and easier to reason about than doing both in one loop.

	// ###Test 3
	// Now try it with []string{"cat", "dog", "cat", "fish", "dog", "cat"} — count all words, delete any that appear only once, and print the remaining counts. Paste your solution when ready!
	animals := []string{"cat", "dog", "cat", "fish", "dog", "cat"}

	counting := make(map[string]int)

	for _, animal := range animals {
		counting[animal]++
	}
	fmt.Println(counting)
	// map[cat:3 dog:2 fish:1]

	for key, value := range counting {
		if value == 1 {
			delete(counting, key)
		}
	}
	fmt.Println(counting)
	// map[cat:3 dog:2]

	//###Example 4
	fmt.Println("//////////////////////////////")
	// This is a groupByLength challenge.
	// The key insight to hold in mind before you code: when you pass a map into a function, you're passing a reference to the same underlying data — so any writes inside the function show up outside it too. No pointers needed.
	// Grouping numbers by whether they are odd or even
	nums := []int{1,2,3,4,5,6}
	groupedNums := groupByEvenOdd(nums)

	for key, values := range groupedNums {
		fmt.Printf("%s: %v\n", key, values)
		// odd: [1 3 5]
		// even: [2 4 6]
	}

	// ###Test 4
	// Now try groupByLength — same idea, but group []string{"go", "map", "key", "range", "for", "len"} by len(word) as the integer key. Paste your solution when ready!
	words := []string{"go", "map", "key", "range", "for", "len", "I"}
	groupedWords := groupByLength(words)

	for key, values := range groupedWords {
		fmt.Printf("%v: %s\n", key, values)
	}
	// Before sorting
	// 5: [range]
	// 1: [I]
	// 2: [go]
	// 3: [map key for len]

	// ADD SORT TO THE KEYS TO PUT IT IN ASCENDING ORDER
	// Example 1 to 3 will help you undertstand this sorting concept better
	fmt.Println("//////////////////////////////")

	keys := make([]int, 0, len(groupedWords))
	// make() is a built-in Go function used to create and initialize slices, maps, and channels.
	// keys := make([]int, 0, len(groupedWords))
	// []int — type: a slice of ints || 0 — starting length (no elements yet) || len(groupedWords) — pre-allocated capacity (avoids resizing as you append)
	for key := range groupedWords {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Loop using sorted keys
	for _, key := range keys {
		fmt.Printf("%v: %s\n", key, groupedWords[key])
	}

	
}

// ####Example 4
func groupByEvenOdd(numbers []int) map[string][]int {
	result := make(map[string][]int)

	for _, num := range numbers {
		if num % 2 == 0 {
			result["even"] = append(result["even"], num) 
		} else {
			result["odd"] = append(result["odd"], num)
		}
	}
	return result
}

// ####Test 4
func groupByLength(words []string) map[int][]string {
	result := make(map[int][]string)

	for _, value := range words {

		length := len(value)
		if length > 0 {
			result[length] = append(result[length], value)	
			
		}
	}
	return result
}