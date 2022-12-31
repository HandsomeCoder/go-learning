package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id      int
	leftCS  *ChopStick
	rightCS *ChopStick
}

func (p Philosopher) eat() {
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("starting to eat %d\n", p.id)
	fmt.Printf("finishing eat %d\n", p.id)
	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

func main() {

	number_philosophers := 5
	chopsticks := make([]*ChopStick, number_philosophers)
	for i := 0; i < number_philosophers; i++ {
		chopsticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, number_philosophers)
	for i := 0; i < number_philosophers; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%number_philosophers]}
	}

	host := make(chan bool, 2)
	var wg sync.WaitGroup

	for i := 0; i < number_philosophers; i++ {

		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			for e := 0; e < 3; e++ {
				host <- true
				philosophers[i].eat()
				<-host
			}
		}()

	}

	wg.Wait()
}
