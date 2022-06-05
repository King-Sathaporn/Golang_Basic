package main

import "fmt"

func main() {
	fnFor(10)
	fnWhile(10)

}

func fnFor(count int) {
	//for loop
	for i := 0; i < count; i++ {
		fmt.Println("Index for loop = ", i)
	}
}

func fnWhile(count int) {
	//Golan dose not support while loop.
	//create while loop with for loop.
	i := 0
	for i < count {
		fmt.Println("Index while loop = ", i)
		i++
		if i >= 5 {
			break // break loop
		}
	}
}
