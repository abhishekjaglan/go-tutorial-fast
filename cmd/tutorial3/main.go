package main

import (
	"fmt"
	"time"
)

func main() {
	//array (fixed size in GO)(multiple declaration styles)

	// var intArr [3]int -> initialization
	var intArr2 [3]int = [3]int{1, 2, 3} // initialization and declaration
	intArr1 := [3]int{1, 2, 3}           // initialization and declaration
	intArr := [...]int{1, 2, 3, 4, 5}    // initilization and declaration
	//intArr[1] = 123
	fmt.Println(intArr[1])
	fmt.Println(intArr[0:3])
	fmt.Println(len(intArr))
	fmt.Println(intArr)
	fmt.Println(intArr1)
	fmt.Println(intArr2)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// slices (wrapper on arrays - no fixed size)(multiple declaration styles)
	var intSlice []int = []int{4, 5, 6, 7}
	fmt.Println("intSlice", intSlice)
	intSlice = append(intSlice, 8)

	intSlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("intSlice2", intSlice2)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	intSlice2 = append(intSlice2, 10)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println("intSlice2 after append", intSlice2)

	var intSlice3 []int = []int{8, 9}
	intSlice = append(intSlice, intSlice3...)
	fmt.Println(intSlice)

	// memory allocation for slice using make(type, len, cap)
	var strSlice []string = make([]string, 2, 4)
	strSlice[0] = "a"
	strSlice[1] = "b"
	strSlice = append(strSlice, "c", "d")
	fmt.Println(strSlice)

	// MAPS - key-value pairs
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Abhishek": 25, "Tanvi": 23}
	fmt.Println(myMap2["Abhishek"])
	fmt.Println(myMap2["Tanvi"])
	fmt.Println(myMap2["Hello"])

	var age, ok = myMap2["Abhishek"] // if value is present in map, true is returned stored in ok
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid arguement to map")
	}

	//delete(myMap2, "Abhishek")
	var abhishek, status = myMap2["Abhishek"]
	if status {
		fmt.Printf("The age is %v\n", abhishek)
	} else {
		fmt.Println("Invalid arguement to map")
	}

	// iterating over maps and slices/arrays
	// FOR loops
	fmt.Println("Iterating over map using for loop\n")
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	fmt.Println("Iterating over array\n")
	for i, v := range intArr {
		fmt.Printf("Value: %v, Index: %v\n", v, i)
	}

	fmt.Println("Iterating over slice\n")
	for i, v := range intSlice2 {
		fmt.Printf("Index: %v	Value: %v\n", i, v)
	}

	fmt.Print("\n")

	for i := 0; i < 5; i++ {
		value := i + 1
		fmt.Printf("Index: %v ---> Value: %v\n", i, value)
	}

	// while loop
	fmt.Println("Iterating using logical while loop")
	var i int = 0

	for i < 5 {
		fmt.Println(i)
		i += 5
	}
	fmt.Printf("Value of i after exiting loop: %v", i)

	i = 0

	for {
		if i > 5 {
			break
		}
		value := i + 1
		fmt.Printf("Index: %v ---> Value: %v\n", i, value)
		i += 5
	}

	// testing time used by slice with and without memory allocation
	var testSliceSize = 1000000
	var testSliceWithoutMemoryAllocation []int8
	var testSliceWithMemoryAllocation []int8 = make([]int8, testSliceSize)

	fmt.Printf("Total time used by slice with %v append operations without memory allocation: %v\n", testSliceSize, testMakeSpeed(testSliceWithoutMemoryAllocation, testSliceSize))
	fmt.Printf("Total time used by slice with %v append operations with memory allocation: %v\n", testSliceSize, testMakeSpeed(testSliceWithMemoryAllocation, testSliceSize))
}

func testMakeSpeed(slice []int8, n int) time.Duration {
	t0 := time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
