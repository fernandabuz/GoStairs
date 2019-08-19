package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string { //função generator
	c := make(chan string)
	go func() { //função anonima
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}() //sempre tem que chamar a função
	return c //recebe os dados a partir da função goroutine
}

func juntar(entrada1, entrada2 <-chan string) <-chan string { //multiplexar usando um select
	c := make(chan string) //aqui eu junto o resultado dos dois canais de entrada em um só
	go func() {            //encapsulo uma chamada de uma função go concorrente
		for { //for infinito sempre selecionando o proximo dado que chega a partir dos canais que eu passei
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
