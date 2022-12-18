package main

import (
	"fmt"
	"time"
)

/*
* In this case we can't predict the value of result because of the race condition
* If condition was handle properly the everytime result value will be 10
 */
func inc(x *int) {
	time.Sleep(10 * time.Millisecond)
	*x++
}

func main() {
	result := 0
	fmt.Println("Initial value of the result variable: ", result)

	for i := 0; i < 10; i++ {
		go inc(&result)
	}

	time.Sleep(time.Second)
	fmt.Println("Finally value of the result variable: ", result)
}
