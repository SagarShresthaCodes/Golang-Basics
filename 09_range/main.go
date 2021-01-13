package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is:", sum)

	//Range with maps
	emails := map[string]string{"Shyle": "925 Pacton Pl", "Keira": "1003 Horeridge Ave"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
		fmt.Println("Address is: " + v)
	}

	for k := range emails {
		fmt.Println("Name is: " + k)
	}
}
