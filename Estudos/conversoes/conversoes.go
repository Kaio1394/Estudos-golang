package main

import (
	f "fmt"
	"strconv"
)

// Operadores lógicos
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2
	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	x, y := 2.4, 2
	result := x / float64(y)
	f.Println("O resultado é: ", result)

	// Convertendo int para str
	f.Println("O resultado é: " + strconv.Itoa(97))

	// Convertendo str para int
	// a função Atoi() retorna dois valores, o resultado da conversão e um erro caso a conversão
	// não seja concluída
	num, _ := strconv.Atoi("123")
	f.Println(num + 20)

	b, _ := strconv.ParseBool("true")
	if b {
		f.Println("Verdadeiro")
	}

	tv50, tv30, sorvete := compras(true, true)
	f.Println(tv50, tv30, sorvete)
}
