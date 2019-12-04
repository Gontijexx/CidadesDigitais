package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"log"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {

	var newUser models.User
	body := r.Body
	bytes, err := util.BodyToBytes(body)
	err = util.BytesToStruct(bytes, &newUser)

	if err = validation.Validator.Struct(newUser); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed) // Status 412
		return
	}

	err = database.InsertNewUser(newUser.Nome, newUser.Email, newUser.Login, newUser.Senha)

	if err != nil {
		log.Println(err)
	}

}
