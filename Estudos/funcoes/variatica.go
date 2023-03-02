package main

import "fmt"

// função com vários parâmetros
func media(numeros ...float64) float64 {
	total := 0.0
	for _, n := range numeros {
		total += n
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println(media(5.0, 2.1, 8.5))
}
