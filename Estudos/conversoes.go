package main

import (
	"fmt"
	"strconv"
)

func main() {
	x, y := 2.4, 2
	result := x / float64(y)
	fmt.Println("O resultado é: ", result)

	// Convertendo int para str
	fmt.Println("O resultado é: " + strconv.Itoa(97))

}
