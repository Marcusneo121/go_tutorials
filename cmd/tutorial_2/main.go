package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 1
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello World" // or `Hello World`
	fmt.Println(myString)

	//some special character outside of ASCII table will have a CHARacter to show length of 2
	//so we need UTF8 package to help for count the length of character
	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a' // RUNE is a data type that stores codes that represent Unicode characters
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var myVar = "text"
	/// same as above =>  myVar := "text"
	fmt.Println(myVar)

	//3 of these below are same
	// var var1, var2 int = 1, 2
	// var var1, var2 = 1, 2
	// var1, var2 := 1, 2

	//constant
	const myConst string = "const value"
	// myConst = "try to change const value which is not working"
	fmt.Println(myConst)
}
