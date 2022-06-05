package main

import "fmt"

func main() {
	value := 5
	if value == 10 {
		fmt.Println("value is 10")
	} else {
		fmt.Println("value is not 10")
	}
	if value > 1 && value < 20 {
		fmt.Println("value is more than 1 and less than 20")
	} else {
		fmt.Println("value is not more than 1 and less than 20")
	}
	if value < 2 || value > 10 {
		fmt.Println("value is more than 1 or less than 20")
	} else {
		fmt.Println("value is not more than 1 or less than 20")
	}
	if value := getValue(); value == "Ok" {
		fmt.Println("Ok")
	} else {
		fmt.Println("-")
	}
	fnSwitchCase()

}

func getValue() string {
	return "Ok"
}

func fnSwitchCase() {
	i := 3
	switch i {
	case 0:
		fmt.Println("i is 0")
		break
	case 1:
		fmt.Println("i is 1")
		break
	case 2:
		fmt.Println("i is 2")
		break
	default:
		fmt.Println("i is not 0, 1 or 2")
		break
	}
}
