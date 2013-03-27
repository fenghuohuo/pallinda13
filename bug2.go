package main

import (
	"fmt"
	"time"
)

// I första versionen av programmet hinner huvud-gorutinen köra klart innan alla ints har skrivits ut, lägger vi till en Sleep() så hinner den köra klart.
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	time.Sleep(100)
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
