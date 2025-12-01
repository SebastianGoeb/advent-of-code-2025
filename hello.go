package main

import (
	"fmt"
	"log"

	"github.com/SebastianGoeb/advent-of-code-2025/days"
)

func main() {
	// configure logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	result, err := days.Day1()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
