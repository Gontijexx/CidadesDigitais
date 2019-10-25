package database

import (
	"database/sql"
	"log"
)

func CheckLogin(login string) (results *sql.Rows, err error) {
	db := ClientSQL()

	results, err = db.Query("SELECT usuario.login FROM usuario WHERE login =?", login)
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT * FROM usuario in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return results, err
}

func CheckSenha(senha string) (results *sql.Rows, err error) {
	db := ClientSQL()

	results, err = db.Query("SELECT usuario.senha FROM usuario WHERE login =?", senha)
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT usuario.senha FROM usuario in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return results, err
}
