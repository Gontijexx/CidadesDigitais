package database

import (
	"CidadesDigitais/models"
	"database/sql"
	"log"
)

// Verifica se o usuario fornecido existe no bando de dados
func CheckLogin(login string) (results *models.Login, err error) {
	db := ClientSQL()
	var cred models.Login

	/*
		QueryRow consulta o bando de dados com o argumento login, se não houver nenhum dado correspondente
		ao argumento passado, QueryRow retorna ErrNoRows. Caso contrário, o retorno eh <nil>. *Row's Scan
		verifica a primeira linha correspondente e descarta o restante.
	*/
	err = db.QueryRow("SELECT usuario.login FROM usuario WHERE login =?", login).Scan(&cred.Login)

	// Tratamento de erro, caso err retorne algo diferente de <nil>
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT usuario.login FROM usuario in database, because: %v\n", err)
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
