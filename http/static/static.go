package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) //passo como parametro o diretório que ele vai ler
	http.Handle("/", fs)                      //quando chegar uma requisição na raiz da minha aplicacao, automaticamente passa por esse handle

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil)) //porta, handle

}

//tem que rodar pelo terminal
//localhost:3000/
//estatico le os arquivos do disco e joga no browser
