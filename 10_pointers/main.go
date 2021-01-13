package main

import "fmt"

func main() {
	a := 5
	b := &a //Assigning b to a pointer of a or the memory address of a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a) //Type of a
	fmt.Printf("%T\n", b) //Type of b

	//Use * to read the value from the address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change the value with pointer
	*b = 10
	fmt.Println(a)
}
