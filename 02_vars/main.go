package main

import "fmt"

func main() {
	//Using var
	//var name = "Brad"

	var age int32 = 37
	const isCool = true

	//Shorthand way of declaring variables, cannot be used outside the function to declare a global variable

	name := "Brad"
	size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
