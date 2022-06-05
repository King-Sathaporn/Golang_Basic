package main

import "fmt"

func main() {
	display()
	display2("Hello World")
	fmt.Println(sum(1, 2))
	fmt.Println(swap(1.2, 3))
	const a, b = 1, 2   //assign multiple value to constant
	x, y := swap2(a, b) //assign to variable
	fmt.Printf("%d,%d => %d,%d", a, b, x, y)

}

func display() {
	//create function with lowercase name is private function
	fmt.Println("Hello World")
}

func display2(message string) {
	//create function with parameter.
	fmt.Println(message)
}

func sum(a, b int) int {
	//create function with return single value.
	return a + b
}

func swap(a float32, b int) (int, float32) {
	//create function with return multiple value.
	return b, a
}

func swap2(a, b int) (x, y int) {
	//create function with return multiple value and assign to variable.
	x = b
	y = a
	return x, y
}
