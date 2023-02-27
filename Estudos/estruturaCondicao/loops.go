package main

import (
	f "fmt"
	"time"
)

type Iterator int64

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
	for {
		f.Println("Executando...")
		time.Sleep(time.Second)
	}

}
