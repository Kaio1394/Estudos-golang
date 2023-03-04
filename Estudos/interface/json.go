package main

import (
	"encoding/json"
	"fmt"
)

// estrutura usada para converter apra json
// `json:""` serve para definir como será os indices no json
// Se esse parâmetro não for definido, o nome de cada indice será  o que foi definido
// como nome da variáveis
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 2800.00, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[1])
}
