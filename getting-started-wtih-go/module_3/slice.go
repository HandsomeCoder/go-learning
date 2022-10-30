// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	arr_slice := make([]int, 0, 3)

	for true {
		var input string

		fmt.Print("Enter integer: ")
		fmt.Scan(&input)

		if input == "X" {
			fmt.Print("Program terminated!")
			break
		}

		intInput, err := strconv.Atoi(input)

		if err == nil {
			arr_slice = append(arr_slice, intInput)
			sort.Ints(arr_slice)
			fmt.Print("Slice Content: ")
			for _, val := range arr_slice {
				fmt.Printf("%d ", val)
			}

		} else {
			fmt.Print("Try Again, Please enter the valid integer")
		}

		fmt.Println()
		fmt.Println()
	}
}
