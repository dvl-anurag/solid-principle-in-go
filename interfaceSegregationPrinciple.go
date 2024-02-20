package main

import "fmt"

// Worker interface with a single method
type Worker interface {
	Work()
}

// Eater interface with a single method
type Eater interface {
	Eat()
}

// Robot implements the Worker interface
type Robot struct{}

// Work method for Robot
func (r Robot) Work() {
	fmt.Println("Robot is working")
}

// Human implements both Worker and Eater interfaces
type Human struct{}

// Work method for Human
func (h Human) Work() {
	fmt.Println("Human is working")
}

// Eat method for Human
func (h Human) Eat() {
	fmt.Println("Human is eating")
}

// performWork takes any type that implements the Worker interface and calls its Work method
func performWork(worker Worker) {
	worker.Work()
}

// performEat takes any type that implements the Eater interface and calls its Eat method
func performEat(eater Eater) {
	eater.Eat()
}
