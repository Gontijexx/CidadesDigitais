package database

import (
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"database/sql"
	"log"
)

//	Verifica se o usuario fornecido existe no bando de dados

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
		return
	}

	defer db.Close()

	return err

}

//	Verifica se a senha fornecida associada ao usuario fornecido existe no banco de dados
//	Funcao OK!

func CheckSenha(login, senha string) (resp bool) {
	db := ClientSQL()
	var senhaDB string

	results, err := db.Query("SELECT usuario.senha FROM usuario WHERE login = ?", login)

	//	Tratamento de erro, caso err retorne algo diferente de <nil>

	if err != nil {
		log.Printf("[WARN] Could not 'SELECT usuario.senha FROM usuario in database, because: %v\n", err)
	}

	//	Converte results em string e atribui a senhaDB
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err := results.Scan(&senhaDB)
		if err != nil {
			log.Printf("[WARN] Could not SCAN in database, because: %v\n", err)
			return
		}
	}

	//	userByte recebe a senha fornecida pelo usuario e a convernte em byte
	userByte := util.PasswordStringToByte(senha)

	//	resp recebe um booleano true se as senhas forem iguais e false caso contrario
	resp = util.ComparePasswords(senhaDB, userByte)

	defer db.Close()

	return resp
}

//	Insere um novo usuario no banco de dados com a senha criptogafada!
//	Funcao OK!

func CreateNewUser(nome, email, login, senha string) (err error) {
	db := ClientSQL()

	//	byteSenha recebe a senha do novo usuario transformada em bytes
	byteSenha := util.PasswordStringToByte(senha)

	//	dbSenha eh a senha criptografada que sera guardada no banco de dados
	dbSenha := util.PasswordByteToHashString(byteSenha)

	//	Exec eh recomendado para ser usado em INSERT e UPDATE
	_, err = db.Exec("INSERT INTO usuario(nome, email, login, senha) VALUES(?,?,?,?)", nome, email, login, dbSenha)

	//	Tratamento do erro, caso o INSERT nao seja feito para err diferente de <nil>
	if err != nil {
		log.Printf("[WARN] Could not 'INSERT newUser in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return err
}

func ReadUser(nome string) (err error) {
	db := ClientSQL()

	var user models.User

	results, err := db.Query("SELECT nome, email, login, senha FROM usuario WHERE nome = ?", nome)

	//	Converte results em string e atribui a senhaDB
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err := results.Scan(&user.Nome, &user.Email, &user.Login, &user.Senha)
		if err != nil {
			log.Printf("[WARN] Could not SCAN in database, because: %v\n", err)
		}
	}

	log.Println(user.Nome, user.Email, user.Login, user.Senha)

	defer db.Close()

	return err
}

func UpdateUser() {
	db := ClientSQL()

	defer db.Close()

	return
}

func DeleteUser(IDUser int) (results sql.Result, err error) {
	db := ClientSQL()

	results, err = db.Exec("DELETE FROM usuario WHERE cod_usuario=?", IDUser)

	log.Println("DELETE!")

	defer db.Close()

	return results, err
}
