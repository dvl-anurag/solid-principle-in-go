package main

// Bird is a superclass
type Bird struct {
	Name string
}

// Flyer is an interface for objects that can fly
type Flyer interface {
	Fly() string
}

// Sparrow is a subclass of Bird and implements the Flyer interface
type Sparrow struct {
	Bird
}

// Fly implements the Flyer interface for Sparrow
func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

// Ostrich is a subclass of Bird but cannot fly
type Ostrich struct {
	Bird
}

// Fly function for Ostrich
func (o Ostrich) Fly() string {
	return "Ostrich cannot fly"
}
