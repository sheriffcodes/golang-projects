package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculator() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		fmt.Println("Example: calculator 10 + 5")
		return
	}

	num1str := os.Args[1]
	operator := os.Args[2]
	num2str := os.Args[3]

	num1, err1 := strconv.Atoi(num1str)
	num2, err2 := strconv.Atoi(num2str)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers")
		return
	}

	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2 //Use "\" to use the muliplication sign in the terminal. Eg  5 \* 2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Number cannot be divided by zero")
			return
		}
		result = num1 / num2 //Use "//" to use the division sign in the terminal. Eg  6 // 2
	default:
		fmt.Println("Error: invalid operator. Use +, -, *, /")
		return
	}

	fmt.Printf("Operator received: '%s'\n", operator)
	fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
	
}