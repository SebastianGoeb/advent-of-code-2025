package main

import (
	"fmt"
	"log"

	"github.com/SebastianGoeb/advent-of-code-2025/greetings"
)

func main() {
	// configure logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	message, err  := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
