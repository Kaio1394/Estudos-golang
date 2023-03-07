package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
func main() {
	c := juntar(
		Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		Titulo("https://www.linkedin.com", "https://www.youtube.com"),
	)

	fmt.Println("1 - ", <-c)
	fmt.Println("2 - ", <-c)
	fmt.Println("3 - ", <-c)
	fmt.Println("4 - ", <-c)
}
