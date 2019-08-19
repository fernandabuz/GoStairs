package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>)", s) //fprintf recebe um writer como parametro
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//tem que rodar pelo terminal
//localhost:3000/horaCerta
//conteúdo dinamico é o que é gerado dinamicamente a partir da execução de um código
