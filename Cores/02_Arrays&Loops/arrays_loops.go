package main

import (
	"fmt"
)

func main() {
	// initializing array
	var arr0 [3]int
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]string = [3]string{"C"}
	arr3 := [3]int{4, 5, 6}
	fmt.Println(len(arr0))
	fmt.Println(arr1, arr2[1], arr3[0:2])
	// infered length of the array
	arr4 := [...]int{1, 2, 3, 4}
	fmt.Println(len(arr4))

	/*
	 *  Slices : Makes array dynamic
	 */
	fmt.Printf("\n ____SLICE_____\n")
	var slice1 []int = []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(slice1, slice2)
	slice3 := append(slice1, slice2...)
	slice3 = append(slice3, 7)
	fmt.Println(slice3, len(slice3), cap(slice3)) // [1 2 3 4 5 6 7] 7 12

	var slice4 []int = make([]int, 3)             // you can initialize a slice using make():
	fmt.Println(slice4, len(slice4), cap(slice4)) // [0 0 0] 3 3
	slice5 := make([]int, 0, 10)                  // make([]Type, length, capacity)
	fmt.Println(slice5, len(slice5), cap(slice5)) // [] 0 10
	index := 9
	if len(slice5) > index {
		// when you initialize a slice with make(), you cant index beyond the length even if the capacity is greater.
		fmt.Println(slice5[index]) // without control flow this will crash
	}
	reslice := slice5[:3]
	fmt.Println(reslice) // [0 0 0] to access the zero vlaues, you must resize first still within the cap
	// append() : returns a new slice with the appended value
	var arr5 []int = []int{1, 2, 3}
	fmt.Println(len(arr5), cap(arr5))
	arr5 = append(arr5, 4)

	/*
	 *  Map : Key value pairs
	 */

	fmt.Printf("\n ____MAP_____\n")
	var ageMap map[string]int = map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	ageMap2 := map[string]int{
		"Charlie": 35,
	}
	ageMap3 := make(map[string]int) // best when populating dynamically

	fmt.Println(ageMap, ageMap2, ageMap3)

	//second value option when fetching a value from a map
	age, ok := ageMap["Charlie"]
	if !ok {
		fmt.Println("Charlie not found")
	} else {
		fmt.Println("Charlie's age:", age)
	}
	delete(ageMap, "Charlie")
	fmt.Println(ageMap)

	// adding a value on a map initalized using make()
	ageMap3["Charlie"] = 35
	ageMap3["Nichi"] = 10
	fmt.Println(ageMap3)

	//LOOPS =======================

	fmt.Printf("\n ____LOOPS_____\n")

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

	fmt.Printf("\n ____LOOPS WITH ARRAYS,SLICES,MAPS_____\n")

	//for array and slicesa
	for index, value := range arr1 { // takes the index, value :=  0 to len - 1. You can omit the index by _
		fmt.Println(index, value)
	}

	for key, value := range ageMap3 { // iterates over the map, similar init array
		fmt.Println(key, value)
	}
}
