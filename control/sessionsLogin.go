package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"log"
	"net/http"
)

// Verificar se os dados para login estao corretos

func Login(w http.ResponseWriter, r *http.Request) {
	/*
		Tratamento dos dados vindos do front-end
		Nesse caso, o request pega login e senha
	*/
	var login models.Credentials
	body := r.Body
	bytes, err := util.BodyToBytes(body)
	err = util.BytesToStruct(bytes, &login)

	// Checks if struct is a valid one
	if err = validation.Validator.Struct(login); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed) // Status 412
		return
	}

	_, err1 := database.CheckLogin(login.Login)
	_, err2 := database.CheckSenha(login.Login, login.Senha)

	if (err1 != nil) || (err2 != nil) {
		w.WriteHeader(http.StatusForbidden) // Status 403
	} else {
		w.WriteHeader(http.StatusAccepted) // Status 202
	}

	return
}
