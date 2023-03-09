package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

// função cria um chanel e armazena o valor de cada operação
func number(x, y int) <-chan int {
	c := make(chan int)
	go func(x, y int) {
		sum := x + y
		c <- sum
	}(x, y)
	go func(x, y int) {
		dif := x - y
		c <- dif
	}(x, y)
	go func(x, y int) {
		mult := x * y
		c <- mult
	}(x, y)

	return c
}

func oMaisRapido(url1, url2, url3 string) string {
	c1, c2, c3 := Titulo(url1), Titulo(url2), Titulo(url3)
	// select -> estrutura específica para goroutine
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.youtube.com",
	)

	fmt.Println(campeao)
	c := number(5, 5)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
