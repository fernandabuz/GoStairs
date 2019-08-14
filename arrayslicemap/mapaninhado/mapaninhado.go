package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{ //chave tipo string e valor tipo map de chave string de valor tipo float64 kk
		"A": {
			"Fernando": 52.0,
			"Taynara":  95.3,
			"Lucas":    65.9,
		},
		"B": {
			"Francisca": 89.9,
			"Thiago":    56.3,
		}, //não esquecer da vírgula
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, numero := range funcs {
			fmt.Println(nome, numero)
		}
	}
}
