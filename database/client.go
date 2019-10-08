package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ClientSQL() {
	log.Println("[START] Go connecting...")

	//definindo db como nosso banco de dados
	//sql.Open tem como parametros nome do drive "mysql", no caso
	//e o segundo parametro eh o endereco de acesso username, password e endereco de IP do localhost e o nome do db
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cidades_digitais_db")
	if err != nil {
		log.Printf("[ERROR] Client SQL not connected because, %v\n", err)
		return
	}

	log.Println("[SUCCESS] Database connected")

	//defer serve para executar esta função de fechamento do banco de dados após todas as outras serem
	//executadas
	defer db.Close()

}
