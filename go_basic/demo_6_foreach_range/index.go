package main

import "fmt"

func main() {

	items := []string{"apple", "banana", "orange"}
	for i, item := range items {
		//i start from 0
		fmt.Println(i+1, item)
	}

	for _, item := range items {
		//if you don't need index, use _
		fmt.Println(item)
	}
}
