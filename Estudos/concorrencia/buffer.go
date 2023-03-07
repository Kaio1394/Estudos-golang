package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	// após o c <- 3 não será executado o código, pois foi definido o tamanho
	// do buffer igual a 3
	fmt.Println("Executou")
	c <- 5
}

func main() {
	// tam buffer igual a 3
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

}
