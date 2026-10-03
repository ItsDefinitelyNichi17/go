# Functions
* Functions in go is first-class citizen
## Basic declaration of a function
```go
func mathAdd(num1, num2 int16) int16 {
	return num1 + num2
}
```
## Multiple and named return values
```go
func mathDiv(num1, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("division by zero") // you can also used fmt.Errorf("division by zero")
		return 0, 0, err
	}
	return num1 / num2, num1 % num2, err
}
```

## Everything is passed by value
```go
func increment(x int) int {
	return x + 1
}

func incrementPointer(x *int) int {
	*x++
	return *x
}
```

## defer statement
Schedules a task to run automatically at the very end of a function—no matter where or how the function exits (even on error). Follows LIFO order.
```go
func deferFunc() {
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")

	fmt.Println("Function body")
}

```

## First Class and Closure
Functions can be stored in variables, passed as arguments, or created anonymously.
```go
func main() {
		num := 
		var sub = mathSub(10)
		fmt.Println(sub(5))
}
func mathSub(x int) func(int) int {
	return func(y int) int {
		return x - y
	}
}
```

## Variadic Functions
Accept zero or more arguments of a specific type. Inside the function, the parameter is treated as a slice. Somewhat similat to spread opeator(...) but not directly equivalent.
```go
func sum(nums ...int) int {
	
}
```
