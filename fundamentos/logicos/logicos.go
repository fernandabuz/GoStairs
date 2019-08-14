package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2    //equivalente ao ou exclusivo
	comprarSorvete := trab1 || trab2 //ou

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
