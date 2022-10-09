// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters 'i', 'a', and 'n'.
// The program should print "Found!" if the entered string starts with the character 'i', ends with the character 'n',
// and contains the character 'a'. The program should print "Not Found!" otherwise.
// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print
// "Found!" for the following example entered strings,"ian", "Ian", "iuiygaygn", "I d skd a efju N".
// The program should print "Not Found!" for the following strings, "ihhhhhn", "ina", "xian".

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter any string: ")

	var str string
	fmt.Scanf("%s%", &str)

	ln := len(str)
	if ln < 3 {
		fmt.Print("Not Found!")
		return
	}

	ln--
	f, l := str[0], str[ln]

	if !(f == 'i' || f == 'I') || !(l == 'n' || l == 'N') {
		fmt.Print("Not Found!")
		return
	}

	for i := 1; i < ln; i++ {
		if str[i] == 'a' {
			fmt.Print("Found!")
			return
		}
	}

	fmt.Print("Not Found!")
}
