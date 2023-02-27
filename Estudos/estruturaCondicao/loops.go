package main

import (
	f "fmt"
	"reflect"
	"time"
)

// Redefinindo o tipo
// Sintaxe semelhante a c++
type Iterator int64

// interface é um tipo genérico
func sum(a, b interface{}) interface{} {
	isInt := reflect.TypeOf(a) == reflect.TypeOf(0) && reflect.TypeOf(b) == reflect.TypeOf(0)
	isFloat := reflect.TypeOf(a) == reflect.TypeOf(0.0) && reflect.TypeOf(b) == reflect.TypeOf(0.0)

	if isInt {
		a1 := reflect.ValueOf(a).Int()
		b1 := reflect.ValueOf(b).Int()
		return a1 + b1
	} else if isFloat {
		a1 := reflect.ValueOf(a).Float()
		b1 := reflect.ValueOf(b).Float()
		return a1 + b1
	}
	return 0
}

func main() {
	var i Iterator = 1

	// em golang não existe o whilte, logo foi criado esse tipo de for para substituí-lo
	for i <= 10 {
		f.Println(i)
		i++
	}

	// for tradicional
	for y := 0; y <= 10; y++ {
		f.Println(y)
	}

	// Lanço infinito
	for {
		f.Println("Executando...")
		time.Sleep(time.Second)
	}

}
