package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		fname string
		lname string
		age   int
	)

	fmt.Print("Enter your first and last name: ")
	n, err := fmt.Scanf("%s %s", &fname, &lname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter your age: ")
	m, err := fmt.Scanf("%d", &age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d items scanned. Hello %s %s!\n", n, fname, lname)
	fmt.Printf("%d items scanned. %d years old!\n", m, age)
}
