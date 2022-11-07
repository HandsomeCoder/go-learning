package main

import (
	"fmt"
	"strconv"
)

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func bubbleSort(arr []int) {
	ln := len(arr)

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func getInt(prompt string) int {
	var num int = -1
	for true {
		var input string

		fmt.Print(prompt)
		fmt.Scanln(&input)

		intInput, err := strconv.Atoi(input)

		if err == nil {
			num = intInput
			break
		} else {
			fmt.Println()
			fmt.Println("Try Again, Please enter the valid integer")
			fmt.Println()
		}
	}

	return num
}

func main() {

	arr_slice := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		num := getInt(fmt.Sprintf("Enter Integer %d: ", i+1))
		arr_slice = append(arr_slice, num)
	}

	fmt.Println("Slice Content: Before Sorting")
	for _, val := range arr_slice {
		fmt.Printf("%d ", val)
	}

	bubbleSort(arr_slice)

	fmt.Println("\n\nSlice Content: After Sorting")
	for _, val := range arr_slice {
		fmt.Printf("%d ", val)
	}
}
