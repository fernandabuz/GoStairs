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
	defer db.Close() //bom fechar sempre depois de executar o erro

	//update
	stmt, _ := db.Prepare("update usuarios set nome=? where id=?")
	stmt.Exec("Fernanda Etc", 1)
	stmt.Exec("Lalala lala", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id=?")
	stmt2.Exec(2)
	stmt2.Exec(4)
}
