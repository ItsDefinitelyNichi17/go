package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/**
	 * Go is a strongly typed language
	 */

	var num3 = 5
	var num uint = 69 // var name type = value
	var str string = "NICHI"
	var float1 float32 = 3.15
	var float2 float64 = 3.15
	var char rune = 'N' // rune is equivalent to char
	var bool bool = true

	num4 := 4                   //short variable declaration, inferred type
	num3, num4, num5 := 3, 4, 5 //variable can be redecalre with this, but there should be a new variable

	fmt.Println(num, str, num3, float1, float2, char, bool, num4, num5)
	fmt.Println(len(str))                    // get the length of the string
	fmt.Println(utf8.RuneCountInString(str)) // get the length of the string utf8
}
