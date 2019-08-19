package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:stairs/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin() //inicializa uma transação e da o objeto "tx" e a partir dele insere os comandos sql
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
