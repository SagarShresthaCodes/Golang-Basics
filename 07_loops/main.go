package main

import "fmt"

func main() {
	//Long method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d fizzfuzz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d fuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d fizz\n", i)
		} else {
			fmt.Println(i)
		}

	}
}
