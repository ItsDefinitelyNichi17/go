package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 5
	var sumOf int = sum(1, 2, 3, 4, 5)
	fmt.Println(sumOf)
	var result int16 = mathAdd(3, 5)
	fmt.Println(result)

	var sub = mathSub(10)
	fmt.Println(sub(5))

	increment1 := increment(5)
	fmt.Println(increment1, num) // 6, 5
	increment2 := incrementPointer(&num)
	fmt.Println(increment2, num) // 6, 6

	deferFunc()
}

// ---name(parameter type) return type{} ---
func mathAdd(num1, num2 int16) int16 {
	return num1 + num2
}

// ---function can return multiple values---
func mathDiv(num1, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("division by zero") // you can also used fmt.Errorf("division by zero")
		return 0, 0, err
	}
	return num1 / num2, num1 % num2, err
}

// ---Variadic Function: Accept zero or more arguments of a specific type, param treated as a slice---
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// ----First class and closure: ----
func mathSub(x int) func(int) int {
	return func(y int) int {
		return x - y
	}
}

// --- pass by value ---
func increment(x int) int {
	return x + 1
}

func incrementPointer(x *int) int {
	*x++
	return *x
}

// --- defer ---
func deferFunc() {
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")

	fmt.Println("Function body")
}
