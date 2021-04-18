package main

import (
	"example/go_challenges/data_types"
	"example/go_challenges/first_steps"
	"example/go_challenges/http_server"
	"example/go_challenges/store"
	"fmt"
)

func main() {
	// testFirstSteps()
	// testDataTypes()
	//testHttpServer()
	testOnlineStore()
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

func testHttpServer() {
	http_server.Run()
}

func testOnlineStore() {
	teo, _ := store.CreateEmployee("Teo", "Gutierrez", 500)

	fmt.Println(fmt.Sprintf("Credits = %f", teo.CheckCredits()))

	if credits, err := teo.AddCredits(250); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	if _, err := teo.RemoveCredits(2500); err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	teo.ChangeName("Amaranto")

	fmt.Println(teo)
}
