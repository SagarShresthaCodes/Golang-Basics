package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string

	//firstName, lastName, gender string
	//age                         int
}

//Greeting method - value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "."
}

//hasBirthday method - pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried method - pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Initializaing person using struct
	person1 := Person{firstName: "Tina", lastName: "Lee", age: 22, gender: "f"}

	//Alternative way of initializing
	person2 := Person{"Leila", "Matthews", 26, "f"}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	fmt.Println(person2.firstName)

	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Ling")
	fmt.Println(person1.greet())
}
