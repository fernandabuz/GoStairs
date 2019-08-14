package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":       5986.56,
		"Gabriela Soares": 8952.63,
		"Willian Palhano": 9865.10, //último elemento tem que ter vírgula também
	}

	funcsESalarios["Fernanda Buzzi"] = 9999.23
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
