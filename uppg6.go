package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	var sum int
	for i := range a {
		sum += a[i]
	}
	res <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	fmt.Println(<-ch + <-ch)
}
