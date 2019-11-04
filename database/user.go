package database

import (
	"CidadesDigitais/models"
	"database/sql"
	"log"
)

func CheckLogin(login string) (err error, l *models.Login) {
	db := ClientSQL()

	log.Print(login)
	var cred models.Login
	err = db.QueryRow("SELECT usuario.login FROM usuario WHERE login =?", login).Scan(&cred.Login)

	defer db.Close()

	return err
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
