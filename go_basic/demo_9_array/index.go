package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var array1 = [5]int{1, 2, 3, 4, 5}
	var array2 = []int{}
	fmt.Println(array1[0])
	fmt.Println(array1)
	fmt.Println(array2)

	for _, item := range array1 {
		fmt.Println(item)
	}
	for i := 0; i < len(array1)*2; i++ {
		array2 = append(array2, i)
	}

	//------------------------------------------------------//

	fmt.Println(array2)

	var sliceArray = make([]int, 5, 10) // 5 is length, 10 is capacity
	showSlice(sliceArray)
	//------------------------------------------------------//
	var numbers []int
	showSlice(numbers)
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
	}
	showSlice(numbers)
	//! Difference between static and dynamic array.
	//! Static array is allocated memory at compile time.
	//! Dynamic array is allocated memory at runtime.
}

func showSlice(array []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array) // cap() is max size of array
}
