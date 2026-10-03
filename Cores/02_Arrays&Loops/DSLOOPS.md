# Data Structures
Check the code on how these works
 ## Array
```go
 // initializing array
	var arr0 [3]int
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]string = [3]string{"C"}
	arr3 := [3]int{4, 5, 6}
```
* Fixed Length: Array length cannot be change or dynamically push an element
* Same Type: Array cannot take multiple types 
* Indexable: can be accessed through index
* Contiguous Array: all the elements are stored side-by-side in one uninterrupted block of physical memory.
## Slice
```go
	var slice1 []int = []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	var slice4 []int = make([]int, 3) // make len cap = 3
	slice5 := make([]int, 0, 10)// make([]Type, length, capacity)
	
````
* Dynamic size, reference-like wrapper around an array. It gives arrays additional functions
* A slice does not actually store data directly. It is a lightweight header containing 3 things:
1. Pointer to the underlying array
2. Length of the slice (len)
3. Capacity of the slice (cap)
* You can Initialize a slice using make(T, len, cap) : predefined allocation of memory 
## Map
```go
var ageMap map[string]int = map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	ageMap2 := map[string]int{
		"Charlie": 35,
	}
	ageMap3 := make(map[string]int) // best when populating dynamically
````
Just like dictionay or other map data structure on other languages, it stores key-value pairs.

# Loops
 `for` is the only keyword used in go for loops.

## The Classic loops
```go
 	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
 	}
  
 	// 0 to N-1
	for i := range 3 { // same as above
		fmt.Println(i)
	}

	n := 0
	for n < 3 { // the while loop equivalent
		n++
		fmt.Println(n)
	}
	
	for { // the while(true) equivalent
		fmt.Println("stopped")
		break
	}
```

## Iterating over slices, maps, arrays
```go
	//for array and slicesa
	for index, value := range arr1 { // takes the index, value :=  0 to len - 1. You can omit the index by _
		fmt.Println(index, value)
	}

	for key, value := range ageMap3 { // iterates over the map, similar init array
		fmt.Println(key, value)
	}

```
* The value of `range` is only a copy of the original value, not a reference to the original value
