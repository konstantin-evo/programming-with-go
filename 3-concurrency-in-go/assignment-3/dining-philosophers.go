package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numMeals        = 3
)

var wg sync.WaitGroup
var host = make(chan struct{}, 2)

type Chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Create philosopher, assign them 2 chopsticks and send them to the dining table
	philosophers := make([]*philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &philosopher{
			id: i, leftChopstick: chopsticks[i], rightChopstick: chopsticks[(i+1)%numPhilosophers]}
		wg.Add(1)
		go philosophers[i].eat()
	}
	wg.Wait()
}

func (p philosopher) eat() {
	defer wg.Done()
	for i := 0; i < numMeals; i++ {
		host <- struct{}{} // capture the slot at the host
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		say("eating", p.id)
		time.Sleep(time.Second)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		<-host // unblocked the chanel after eating

		say("finished eating", p.id)
		time.Sleep(time.Second)
	}
}

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}
