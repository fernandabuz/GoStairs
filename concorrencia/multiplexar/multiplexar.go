package main

import (
	"fmt"

	"github.com/fernandabuz/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem //encaminhando o dado do origem pro destino - quando não tiver dados na origem, o for para e aguarda
	}
}

//multiplexar - misturar (mensagens) de fontes diferentes num mesmo canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c) //entrada1 origem e c destino
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("http://stairs.studio", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
