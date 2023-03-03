package main

import (
	"fmt"
	"reflect"
)

// função anônima
var soma = func(a, b int) int {
	return a + b
}

// passando um ponteiro como parâmetro
func change(item *int) {
	*item += 2
}

func multiplicacao(a, b interface{}) interface{} {
	return reflect.ValueOf(a).Int() * reflect.ValueOf(b).Int()
}

func soma2(a, b interface{}) interface{} {
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

// passando parâmetros genéricos e retornando valor também genérico
func teste(funcao func(interface{}, interface{}) interface{},
	a, b interface{}) interface{} {
	isInt := reflect.TypeOf(a) == reflect.TypeOf(0) && reflect.TypeOf(b) == reflect.TypeOf(0)
	isFloat := reflect.TypeOf(a) == reflect.TypeOf(0.0) && reflect.TypeOf(b) == reflect.TypeOf(0.0)

	if isInt {
		a = reflect.ValueOf(a).Int()
		b = reflect.ValueOf(b).Int()
		return funcao(a, b)
	} else if isFloat {
		a = reflect.ValueOf(a).Float()
		b = reflect.ValueOf(b).Float()
		return funcao(a, b)
	}
	return 0
}

func main() {
	fmt.Println(teste(soma2, 1, 2))
	x := 5
	var y *int = &x
	change(y)

	// fmt.Println(x)
	fmt.Println(*y)
}
