package main

import (
	"fmt"
)

func main() {
	//map have key and value
	var colors = map[string]string{
		"red":   "#ff0000", //key:value
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])    //map can get value by key
	fmt.Println(colors["yellow"]) //if key not exist, return ""
	delete(colors, "red")         //delete key
	fmt.Println(colors)           // if you delete key, value will be ""
	//--------------------------------------------------------------------------------
	var number = make(map[string]map[string]int) //create map multi level
	fmt.Println(number)
	number["item_name_1"] = make(map[string]int)
	fmt.Println(number)
	number["item_name_1"]["id"] = 1234
	fmt.Println(number)
	number["item_name_1"]["price"] = 20
	fmt.Println(number)
	number["item_name_2"] = map[string]int{"id": 9876, "price": 15}
	fmt.Println(number)
	fmt.Println(number["item_name_1"])
	fmt.Println(number["item_name_1"]["id"])
}
