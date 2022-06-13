package singleton

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	once     sync.Once
	instance *greeter
)

type Greeter interface {
	Greet() string
}

type greeter struct {
	greeting string
}

func create() *greeter {
	g := &greeter{
		greeting: fmt.Sprintf("Hello from Greeter - %d", rand.Intn(100)),
	}

	return g
}

func New() Greeter {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating greeter instance now.")
				instance = create()
			})
	} else {
		fmt.Println("Greeter instance already created.")
	}

	return instance
}

func (g *greeter) Greet() string {
	return g.greeting
}
