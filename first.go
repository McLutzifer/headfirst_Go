package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("First argument", "Second argument")
	fmt.Println("Another argument")
	fmt.Println(strings.Title("head first go"))
	fmt.Println(math.Floor(2.75))
	fmt.Println(7 / 5.0)

	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go"))
	fmt.Println(reflect.TypeOf('A'))

	var test string = string(1174)
	fmt.Println(test)
}
