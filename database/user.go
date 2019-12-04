package database

import (
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"log"
)

// Verifica se o usuario fornecido existe no bando de dados

func CheckLogin(login string) (err error) {
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

	return err

}

// Verifica se a senha fornecida associada ao usuario fornecido existe no banco de dados

func CheckSenha(login, senha string) (err error) {
	db := ClientSQL()
	var cred models.Credentials

	err = db.QueryRow("SELECT usuario.senha FROM usuario WHERE login = ? and senha = ?", login, senha).Scan(&cred.Senha)

	// Tratamento de erro, caso err retorne algo diferente de <nil>

	if err != nil {
		log.Printf("[WARN] Could not 'SELECT usuario.senha FROM usuario in database, because: %v\n", err)
	}

	defer db.Close()

	return err
}

func InsertNewUser(nome, email, login, senha string) (err error) {
	db := ClientSQL()

	byteSenha := util.PasswordStringToByte(senha)
	dbSenha := util.PasswordByteToHashString(byteSenha)

	len := len(dbSenha)

	log.Println(len)

	_, err = db.Query("INSERT INTO usuario(nome, email, login, senha) VALUES(?,?,?,?)", nome, email, login, dbSenha)

	if err != nil {
		//we need to change the menssage log
		log.Printf("[WARN] Could not 'INSERT newUser FROM usuario in database, because: %v\n", err)
	}

	defer db.Close()

	return err
}
