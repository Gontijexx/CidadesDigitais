package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"log"
	"net/http"
)

// Criar novo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser models.User
	body := r.Body
	bytes, err := util.BodyToBytes(body)
	err = util.BytesToStruct(bytes, &newUser)

	// Verifica se os dados que chegaram correspodem as pre-condicoes

	if err = validation.Validator.Struct(newUser); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed) // Status 412
		return
	}

	// Verifica se o metodo recebido eh o esperado

	if r.Method == "POST" {

		// Consulta no bando de dados se o login ja existe
		err = database.CheckLogin(newUser.Login)

		/*
			Se ja existir o login no banco de dados, retorna um erro
			Senao o novo usuario é criado
		*/
		if err == nil {
			log.Printf("[WARN] invalid user information, because, %v\n", err)
			w.WriteHeader(http.StatusPreconditionFailed) // Status 412
			return
		} else {
			database.InsertNewUser(newUser.Nome, newUser.Email, newUser.Login, newUser.Senha)
		}
	}

}
