package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(3 * time.Second)
	channel <- 4 * base
}

func main() {
	channel := make(chan int)
	go doisTresQuatroVezes(2, channel)
	fmt.Println("A")

	a, b := <-channel, <-channel
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-channel)
}
