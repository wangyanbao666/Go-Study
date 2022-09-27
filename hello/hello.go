package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
