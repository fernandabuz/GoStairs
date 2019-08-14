package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.") //não cria uma nova linha

	fmt.Println(" Nova") //aqui quebra a linha
	fmt.Println("linha.")

	x := 3.141516
	//fmt.Println("O valor de x é " + x) não funciona colocando um string + float

	xs := fmt.Sprint(x)                 //invés de imprimir no console, vai retornar um string com essa função
	fmt.Println("O valor de x é " + xs) //aqui imprime tudo string

	//OU
	fmt.Println("O valor de x é", x, "!!!!!")
	fmt.Printf("O valor de x é %.2f", x) //%.2f quer dizer que eu quero que imprima apenas 2 casas decimais
	//o %f quer dizer que ele vai imprimir um tipo float, se fosse %s seria string

	a := 1
	b := 2.5555
	c := false
	d := "eita"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) //printf fica na mesma linha pra quebrar usa \n
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
