package main

import (
	"fmt"
)

// passando map como parâmetro de uma função
func add(estrutura map[int]string, cpf int, name string) {
	estrutura[cpf] = name
}

// deletando um elemento do map
func deleteElement(map_ map[int]string, key int) {
	delete(map_, key)
}

func main() {
	// criando um map
	aprovados := make(map[int]string)
	aprovados[1241415] = "Kaio"
	aprovados[5456465] = "Eliziane"
	aprovados[56164516] = "Maria"
	aprovados[651465451] = "Jorge"
	aprovados[54545454] = "Yuri"
	fmt.Println(aprovados)

	add(aprovados, 154151541, "José")
	for key, value := range aprovados {
		fmt.Println(key, value)
	}

	// inicializando map com valores
	// obs.: sempre o último eleento deve termianr com ','
	reprovados := map[int]string{
		65151651:      "teste",
		5461156561151: "T",
	}
	fmt.Println(reprovados)
	deleteElement(reprovados, 65151651)
	fmt.Println(reprovados)

	// maps aninhados
	funcPorLetra := map[string]map[string]float64{
		"K": {
			"Kaio Cândido Santiago": 4200.00,
			"Kevyn":                 5000.00,
		},
	}
	fmt.Println(funcPorLetra["K"]["Kaio Cândido Santiago"])

}
