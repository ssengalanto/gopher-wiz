package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct {
}

var instance *singleton

func newInstance() *singleton {
	return &singleton{}
}

func Instance() *singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = newInstance()
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
