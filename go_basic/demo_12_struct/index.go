package main

import (
	"fmt"
)

func main() {
	showItem("iPhone", 1000, 10)
	var x product
	x.name = "iPhone"
	x.price = 35000
	x.stock = 50
	fmt.Println(x)
	showProduct(x)
	update(&x)
	showProduct(x)
	y := x.clear()
	fmt.Println(y)
	z := x.setDiscount(10)
	fmt.Println(z)

}

func showItem(name string, price int, stock int) {
	fmt.Printf("%s %d %d\n", name, price, stock)
}

type product struct {
	//! Golang does not support class.
	//! But we can use struct to simulate class.
	name  string
	price int
	stock int
}

func showProduct(x product) {
	fmt.Printf("%s %d %d\n", x.name, x.price, x.stock)
}

//-----------------------------------------------------------------------------

func update(x *product) {
	x.price = x.price + 10000
	x.stock = x.stock - 10
}

//-----------------------------------------------------------------------------

func (x product) clear() product {
	x.name = "-"
	x.price = 0
	x.stock = 0
	return x
}

func (x product) setDiscount(discount int) product {
	x.price = x.price - discount
	return x
}