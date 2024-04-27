package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello World 123"
	printMe(printValue)

	var division, remainder, err = intDivision(11, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The result of the integer division is %v with the remainder %v", division, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
