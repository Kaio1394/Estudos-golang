package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		primo := inicio
		for {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
			primo++
		}
	}
	close(c)
}
func main() {
	ch := make(chan int, 30)

	go primos(60, ch)

	for primo := range ch {
		fmt.Printf(" %d ", primo)
	}
}
