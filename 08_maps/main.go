package main

import "fmt"

func main() {
	//Define map - map with a string key declared in big brackets and a string value
	//emails := make(map[string]string)

	//Assign key values
	//emails["Bob"] = "bob@gmail.com"
	//emails["Shannon"] = "shannon@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	//Shortcut way of declaring a map and adding key-value pairs
	emails := map[string]string{"Bob": "bob@gmail.com", "Shannon": "shannon@gmail.com", "Mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//Deleting a Key value pair from the map
	delete(emails, "Bob")
	fmt.Println(emails)
}
