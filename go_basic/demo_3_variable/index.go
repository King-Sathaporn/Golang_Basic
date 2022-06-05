package main

import "fmt"

func main() {
	fmt.Println("Start")
	var tmp1 int = 1         // var can be changed value
	const tmp2 string = "Hi" //! const can't be changed
	var tmp3 bool = true
	var tmp4 float32 = 1.5
	fmt.Println(tmp1, tmp2, tmp3, tmp4)

	tmp3 = false
	fmt.Println(tmp3)

	tmp1 = int(tmp4 / 2)
	fmt.Println(tmp1) //? if answer type is int and value is less than 1.0, it will be 0

	var tmp5 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(tmp5[i])
	}

	var count1 int = 0
	count1++ //? count1++ = count1 + 1
	print(count1)
}
