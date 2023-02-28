package main

import (
	"fmt"
)

func main() {
	// declaração de um array
	var notas [3]float64
	var ptr *float64 = &notas[0]
	fmt.Println(ptr)

	notas[0], notas[1], notas[2] = 2.2, 4.5, 8.8

	sum := 0.0
	for i := 0; i < len(notas); i++ {
		sum += notas[i]
	}
	fmt.Println(sum / float64(len(notas)))
}
