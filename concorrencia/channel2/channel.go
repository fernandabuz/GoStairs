package main

import (
	"fmt"
	"time"
)

//channel - é a forma de comunicação entre as goroutines
//channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) //goroutin causa assincronismo e o chanal é o que sincroniza
	fmt.Println("A")

	a, b := <-c, <-c //recebendo os dados do canal
	fmt.Println("B")

	fmt.Println(a, b)
	fmt.Println(<-c)
	//fmt.Println(<-c) não tem nada vindo por isso daria erro
}
