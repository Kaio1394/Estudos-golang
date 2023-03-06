package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// informa para o programa a quantidade de goroutines
	wg.Add(2)

	go func1()
	go func2()

	// Especifica a quantidade de goroutine que o programa tem que executar antes
	// de finalizar a função main
	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1: ", i)
		time.Sleep(time.Second)
	}
	// Fim da goroutine
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2: ", i)
		time.Sleep(time.Second)
	}
	// Fim da goroutine
	wg.Done()
}
