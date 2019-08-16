package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/*
		fale("Maria", "Por que você não fala comigo?", 3) são funções seriais, primeiro executa uma e quando terminar executa a outra
		fale("Pedro", "Só posso falar depois de você", 1)
	*/

	/*
		go fale("Maria", "Hey", 500)
		go fale("Pedro", "Oi..", 500)
		fmt.Println("Fim!") como a main terminou antes da funcao fale poder imprimir algo (porque demora 1s) não teve output de fale

		time.Sleep(time.Second * 5) se colocar isso antes do "Fim" ele executa o fale 4 segundos depois termina
	*/

	/*
		threads são linhas de execução independentes - em go são goroutines
		goroutine só continua executando se a thread principal do seu programa estiver executando
	*/

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5) //varia as execuções porque são funções rodando independentes e termina em 5 - definida no fale sem go
}
