package database

import (
	"CidadesDigitais/models"
	"log"
)

// Verifica se o usuario fornecido existe no bando de dados

func CheckLogin(login string) (results *models.Credentials, err error) {
	db := ClientSQL()
	var cred models.Credentials

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

func CheckSenha(login, senha string) (results *models.Credentials, err error) {
	db := ClientSQL()
	var cred models.Credentials

	log.Print(senha)

	err = db.QueryRow("SELECT usuario.senha FROM usuario WHERE login = ? and senha = ?", login, senha).Scan(&cred.Senha)

	// Tratamento de erro, caso err retorne algo diferente de <nil>
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT usuario.senha FROM usuario in database, because: %v\n", err)
	}

	defer db.Close()

	return results, err
}
