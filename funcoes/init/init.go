package main

import "fmt"

func init() { //mesmo sem chamar na main, a init é executada antes
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
