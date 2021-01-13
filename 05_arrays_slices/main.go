package main

import "fmt"

func main() {
	//Arrays - fruitArr will have 2 string values
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declaring and Assigning values short way
	fruitArr := [2]string{"Apple", "Orange"}
	fruitArrSlice := []string{"Apple", "Orange", "Mango"} //array slice does not need a specific number

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArrSlice)
	fmt.Println(len(fruitArr))
	fmt.Println(len(fruitArrSlice)) //length
	fmt.Println(fruitArrSlice[1:3]) //Range

}
