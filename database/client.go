package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //	database driver
)

//	ClientSQL definindo db como nosso banco de dados
func ClientSQL() (db *sql.DB) {
	log.Println("[START] Go connecting...")

	/*
		definindo db como nosso banco de dados sql.Open
		tem como parametros nome do drive "mysql", no caso
		e o segundo parametro eh dataSourceName:
		"username:password@endereco_de_IP_local[tcp(127.0.0.1:3306)]/nome_db"
	*/
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cidades_digitais_db")
	if err != nil {
		log.Printf("[ERROR] Client SQL not connected because, %v\n", err)
		return
	}

	log.Println("[SUCCESS] Database connected")

	return db
}
