package database

import (
	"database/sql"
	"log"
)

func CheckLogin(login string) (results *sql.Rows, err error, db *sql.DB) {
	db = ClientSQL()

	results, err = db.Query("SELECT * FROM usuario WHERE login =?", login)
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT * FROM usuario in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return results, err, db
}

func CheckSenha(senha string) (results *sql.Rows, err error, db *sql.DB) {
	db = ClientSQL()

	results, err = db.Query("SELECT * FROM senha WHERE login =?", senha)
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT * FROM senha in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return results, err, db
}