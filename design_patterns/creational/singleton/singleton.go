package main

import (
	"fmt"
	"log"
	"sync"
)

var once sync.Once

type Singleton struct {
}

var instance *Singleton

func Instance() *Singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = &Singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}

func main() {
	for i := 0; i < 7; i++ {
		go Instance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
