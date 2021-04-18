package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzBuzzUsingSwitch()
	println("Square root is: ", fmt.Sprintf("%f", guestSquareRoot(25, 1, 0)))
	paincNegative()
}

func fizzBuzz(number int) string {
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func fizzBuzzUsingSwitch() {
	for i := 1; i <= 100; i++ {
		println(fizzBuzz(i))
	}
}

func guestSquareRoot(x float64, currguess float64, n int) float64 {
	if n < 10 {
		nextguess := currguess - (currguess*currguess-x)/(2*currguess)
		if nextguess != currguess {
			println("A guess for square root is ", fmt.Sprintf("%f", nextguess))
			return guestSquareRoot(x, nextguess, n+1)
		}
	}
	return currguess
}

func paincNegative() {
	var val int
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("Negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
