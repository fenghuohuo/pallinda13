package main

import (
	"fmt"
	"time"
)

func Remind(text string, paus time.Duration) {
	var current string
	for {
		time.Sleep(paus)
		current = time.Now().Format("15:04")
		fmt.Println(text, current)
	}
}

func main() {
	go Remind("Klockan är %s Dags att äta", time.Second*3)
	go Remind("Klockan är %s Dags att arbeta", time.Second*8)
	Remind("Klockan är %s Dags att sova", time.Second*24)
	select {}
}
