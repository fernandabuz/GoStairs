package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem sempre ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[15469852] = "Pedro"
	aprovados[85697412] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { //primeiro valor definido é sempre a chave do map
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 15469852)
	fmt.Println(aprovados[15469852])
}
