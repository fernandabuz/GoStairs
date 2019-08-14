package main

import "fmt"

func f1() {
	fmt.Println("F1!")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("e F4: %s, %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

//função pura é quando recebe um conjunto de parametros e de forma deterministica sempre gera o mesmo resultado de acordo com o valor dos parametros
//não mexe em nada externo, não tem efeitos colaterais em nada que está fora da função

func main() {
	f1()
	f2("oi", "tchau")
	r3, r4 := f3(), f4("eita", "opa")
	fmt.Println(r3, r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
