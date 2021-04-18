package main

import (
	"errors"
	"fmt"
)

func main() {
	var sequence int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&sequence)
	if res, err := fibonacci(sequence); err == nil {
		fmt.Println(res)
	}

	var romanNumber string
	fmt.Print("Enter a roman number: ")
	fmt.Scanln(&romanNumber)
	if res, err := romanNumeralsTranslator(romanNumber); err == nil {
		fmt.Println(res)
	}
}

func fibonacci(sequence int) ([]int, error) {
	if sequence <= 2 {
		return nil, errors.New("Sequence must be greater than 2")
	}

	serie := make([]int, sequence)
	serie[0], serie[1] = 1, 1

	for i := 2; i < sequence; i++ {
		serie[i] = serie[i-1] + serie[i-2]
	}

	return serie, nil
}

func romanNumeralsTranslator(romanNumber string) (int, error) {
	dict := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	translated := make([]int, len(romanNumber)+1)
	for idx, letter := range romanNumber {
		if val, found := dict[letter]; found {
			translated[idx] = val
		} else {
			fmt.Printf("Letter %c not found!", letter)
		}
	}

	var res int
	for idx := 0; idx < len(romanNumber); idx++ {
		if translated[idx] < translated[idx+1] {
			translated[idx] *= -1
		}
		res += translated[idx]
	}

	return res, nil
}
