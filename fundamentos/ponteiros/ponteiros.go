package main

import "fmt"

func main() {
	i := 1
	var p *int //= nil
	/*
		O asterisco na frente do dado indica que é um ponteiro
		Tem um espaço de memória que armazena um ponteiro e dentro dele tem um endereço de memória que guarda um dado
		Aloquei um espaço de memória que é do tipo ponteiro que dentro dele tem um endereço de memória que está vazio
	*/

	p = &i //p agora guarda o endereco de memoria da variavel i
	*p++   //desreferenciando o ponteiro pra pegar o valor
	i++

	//p++ não posso fazer porque é uma operação aritmética em cima de um ponteiro
	//go não tem aritmética de ponteiros

	fmt.Println(p, *p, i, &i)
}
