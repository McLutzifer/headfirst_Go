package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	customerName = "Damon Cole"
	length, width = 1.2, 2.4

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
}
