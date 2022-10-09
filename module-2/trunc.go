package main

import (
	"fmt"
)

func main() {
	var fnum float32
	fmt.Print("Enter any Floating Point number: ")
	fmt.Scan(&fnum)
	fmt.Printf("Decimal number of entered Floating Point number: %d", int32(fnum))
}
