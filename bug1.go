package main

import (
	"fmt"
	"time"
)

// Eftersom kanaler blockar innan de kan ta emot/skicka ett värde så funkar inte första versionen av programmet.
func main() {
	ch := make(chan string)
	go Print(ch)
	ch <- "Hello world!"
	time.Sleep(100)
}

func Print(ch <-chan string) {
	fmt.Println(<-ch)
}
