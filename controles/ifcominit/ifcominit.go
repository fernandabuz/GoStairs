package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int { //gerando um numero aleatorio
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //também é suportado no switch
		fmt.Println("Ganhou :D")
	}
	fmt.Println("Perdeu :(")
}
