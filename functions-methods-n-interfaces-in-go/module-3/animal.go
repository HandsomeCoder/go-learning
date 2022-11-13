package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	spoken     string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.spoken
}

func getString(prompt string) (string, string) {

	name := ""
	attribute := ""
	for true {
		fmt.Print(prompt)
		fmt.Scanln(&name, &attribute)

		if name != "" && attribute != "" {
			break
		} else {
			fmt.Printf("Try Again, Please enter valid name (cow, bird, or snake) and attribute (eat, move, or speak)\n\n")
		}
	}

	return name, attribute
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	mapping := map[string]Animal{"cow": cow, "bird": bird, "snake": snake}

	for true {
		name, attribute := getString("> ")

		if animal, ok := mapping[name]; ok {
			switch attribute {
			case "eat":
				fmt.Printf("%s %s = %s\n\n", name, attribute, animal.Eat())
			case "move":
				fmt.Printf("%s %s = %s\n\n", name, attribute, animal.Move())
			case "speak":
				fmt.Printf("%s %s = %s\n\n", name, attribute, animal.Speak())
			default:
				fmt.Printf("Try Again, Please enter valid name (cow, bird, or snake) and attribute (eat, move, or speak)\n\n")
			}

		} else {
			fmt.Printf("Try Again, Please enter valid name (cow, bird, or snake) and attribute (eat, move, or speak)\n\n")
		}
	}

}
