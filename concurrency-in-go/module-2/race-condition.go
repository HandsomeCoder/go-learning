package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* In this case we can predent the value of result because of the race condition
* If condition was handle properly the everytime result value will be 400
 */
func main() {
	result := 0
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 200; j++ {
				result++
				time.Sleep(2 * time.Millisecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Finally a becomes:", result)
}
