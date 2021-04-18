package main

import (
	"example/go_challenges/data_types"
	"example/go_challenges/first_steps"
	"fmt"
)

func main() {
	// testFirstSteps()
	// testDataTypes()
}

func testFirstSteps() {
	fmt.Println("Hello World")

	first_steps.FizzBuzzUsingSwitch()
	println("Square root is: ", fmt.Sprintf("%f", first_steps.GuestSquareRoot(25, 1, 0)))
	first_steps.PaincNegative()
}

func testDataTypes() {
	var sequence int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&sequence)
	if res, err := data_types.Fibonacci(sequence); err == nil {
		fmt.Println(res)
	}

	var romanNumber string
	fmt.Print("Enter a roman number: ")
	fmt.Scanln(&romanNumber)
	if res, err := data_types.RomanNumeralsTranslator(romanNumber); err == nil {
		fmt.Println(res)
	}
}
