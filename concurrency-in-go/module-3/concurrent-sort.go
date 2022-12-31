package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func printSlice(title string, slice []int) {
	fmt.Println(title)
	for _, val := range slice {
		fmt.Printf("%d ", val)
	}
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

func getIntegerSlice(prompt string) []int {

	numbers := ""
	elements := make([]int, 0, 4)
	valid := true

	for true {
		valid = true
		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numbers = scanner.Text()
		}

		if numbers == "" {
			fmt.Printf("Try Again, Please enter only number with space separated \n\n")
			continue
		}

		parts := strings.Split(numbers, " ")

		if len(parts) < 4 {
			fmt.Printf("Try Again, Please atleast 4 number with space separated \n\n")
			numbers = ""
			continue
		}

		for _, itr := range parts {
			intInput, err := strconv.Atoi(itr)
			if err == nil {
				elements = append(elements, intInput)
			} else {
				valid = false
				numbers = ""
				elements = []int{}
				fmt.Printf("Try Again, Please enter only number with space separated \n\n")
				break
			}
		}

		if valid {
			break
		}

	}

	return elements
}

func runner(name string, slice []int) {
	printSlice(fmt.Sprintf("From Runner %s: Before Sorting...", name), slice)
	bubbleSort(slice)
	printSlice(fmt.Sprintf("\nFrom Runner %s: After Sorting...", name), slice)
	fmt.Printf("\n\n")
}

func main() {

	threads := 4
	input := getIntegerSlice(fmt.Sprintf("Enter Integers (space separated): "))

	ln := len(input)
	sub_ln := ln / threads

	channel := make(chan []int)

	for i := 0; i < threads; i++ {
		start := i * sub_ln
		end := start + sub_ln

		if i == 3 {
			end = ln
		}

		i := i
		go func() {
			sub_nums := input[start:end]
			runner(strconv.Itoa(i+1), sub_nums)
			channel <- sub_nums
		}()

	}

	merged := make([]int, 0)
	for i := 0; i < threads; i++ {
		sub := <-channel
		merged = append(merged, sub...)
	}

	printSlice("Main Thread: Before Sorting...", merged)
	bubbleSort(merged)
	printSlice("\nMain Thread: After Sorting...", merged)

}
