package main

import "fmt"

func main() {
	// channel
	ch := make(chan int, 1)

	// enviando dados
	ch <- 1
	fmt.Println(<-ch)
}
