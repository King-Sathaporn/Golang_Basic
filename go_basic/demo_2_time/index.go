package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world.")
	time.Sleep(5 * time.Second) //?delay n*1 second
	fmt.Println("END")
}
