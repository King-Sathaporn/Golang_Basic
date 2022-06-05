package main

import "fmt"

func main() {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	showSlice(array)
	array = array[3:] // array = array[start:end] default start = 0 end = len(array)
	showSlice(array)

	var number = []int{0, 1, 2, 3, 4}
	number = removeIndex(number, 2)
	showSlice(number)

}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	//fmt.Println("len=%d cap=%d slice=%v", len(x), cap(x), x)
	//! Println is not supported %d, %v, %T...
}

//?delete specific index function
func removeIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...) // ... is spread operator
}
