package main

import "fmt"

//channel é um tipo da linguagem e pode trafegar dado de dederminados tipos
//channel é um mecanismo de sincronismo - consigo esperar os dados chegarem das funções independentes e só continua se chegar no canal o que ele tá esperando

func main() {
	ch := make(chan int, 1) //o 1 é o buffer

	ch <- 1 //enviando para o canal o valor 1 (escrita)
	<-ch    //saindo do canal determinado valor (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
