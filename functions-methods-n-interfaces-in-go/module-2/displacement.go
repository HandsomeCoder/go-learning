package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(acceleration, velocity, initDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (acceleration*time*time)/2 + (velocity * time) + initDisplacement
	}
}

func getFloat(prompt string) float64 {
	var num float64 = -1
	for true {
		var input string

		fmt.Print(prompt)
		fmt.Scanln(&input)

		intInput, err := strconv.ParseFloat(input, 64)

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

	acceleration := getFloat("Enter Acceleration		: ")
	velocity := getFloat("Enter Initial Velocity		: ")
	initDisplacement := getFloat("Enter Initial Displacment	: ")

	displacementCalculator := GenDisplaceFn(acceleration, velocity, initDisplacement)

	fmt.Print("\n============================================\n")
	time := getFloat("Enter Time			: ")
	fmt.Print("============================================\n\n")

	fmt.Printf("Displacement after time %0.2f : %.2f", time, displacementCalculator(time))
}
