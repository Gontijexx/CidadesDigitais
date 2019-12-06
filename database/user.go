package database

import (
	"CidadesDigitais/models"
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

	/*
		Tratamento de erro, caso err retorne algo diferente de <nil>
		Excluir ou tratar este err de outra forma, CheckLogin esta sendo usada em duas situacoes
		em que para cada uma temos um retorno diferente para os err
	*/
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

// Insere um novo usuario no banco de dados
// Por enquanto a criptografia de senha nao esta funcionando, o banco de dados nao esta aceitando o tamanha da string criptograda

func InsertNewUser(nome, email, login, senha string) (err error) {
	db := ClientSQL()

	/*
		Essas funcoes serao usadas para a criptografia, por enquanto o banco nao esta aceitando o tamanho da string

		byteSenha := util.PasswordStringToByte(senha)
		dbSenha := util.PasswordByteToHashString(byteSenha)

		len := len(dbSenha)

		log.Println(len)
	*/

	// Exec eh recomendado para ser usado em INSERT e UPDATE
	_, err = db.Exec("INSERT INTO usuario(nome, email, login, senha) VALUES(?,?,?,?)", nome, email, login, senha)

	if err != nil {
		//we need to change the menssage log
		log.Printf("[WARN] Could not 'INSERT newUser FROM usuario in database, because: %v\n", err)
	}

	defer db.Close()

	return err
}

func DeleteUser(IDUser int) {

	db := ClientSQL()

	db.Exec("DELETE FROM usuario WHERE cod_usuario=?", IDUser)

	defer db.Close()
}
