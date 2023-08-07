package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func validateInput(input string) error {
	inputType := regexp.MustCompile(`^[01]+$`).MatchString(input)
	inputLen := len([]rune(input))

	if !inputType {
		return errors.New("Input must be 0 or 1.")
	}

	if inputLen > 8 {
		return errors.New("Your input exceeded the character limit.")
	}

	return nil
}

func binaryToDecimal(input string) (int, error) {
	decimalNum := 0

	if inputErr := validateInput(input); inputErr != nil {
		return decimalNum, inputErr
	}

	var remainder int
	index := 0
	num, _ := strconv.Atoi(input)

	for num != 0 {
		remainder = num % 10
		num = num / 10

		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}

	return decimalNum, nil
}

func main() {
	var input string

	fmt.Println(" ---===--- Bin2Dec ---===--- ")
	fmt.Println("This program will convert binary number to decimal. \n ")
	fmt.Println("Please input binary number (Maks. 8 character) : ")
	fmt.Scanln(&input)

	if result, err := binaryToDecimal(input); err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println(" \nReseult : ")
		fmt.Println(result)
	}
}
