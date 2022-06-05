package main

import "fmt"

//!pointer is a variable that holds the memory address of another variable
func main() {
	message := "Hello, World!"
	var messagePointer *string = &message //?create a pointer to the message string
	fmt.Println(message)
	fmt.Println(messagePointer)
	fmt.Println(*messagePointer)        //?dereference the pointer to get the value
	*messagePointer = "Goodbye, World!" //?change the value with the pointer
	fmt.Println(message)
	changeMessage(&message, "New Message 1")
	fmt.Println(message)
	changeMessage(messagePointer, "New Message 2")
	fmt.Println(message)
}

func changeMessage(messagePointer *string, newMessage string) {
	*messagePointer = newMessage //?change the value with the pointer
}
