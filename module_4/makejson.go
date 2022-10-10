// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Address: ")
	fmt.Scanln(&address)

	person := make(map[string]string)

	person["name"] = name
	person["address"] = address

	jstring, _ := json.Marshal(person)

	fmt.Printf("Details: %s", jstring)
}
