package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"log"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var delete models.User
	body := r.Body
	bytes, err := util.BodyToBytes(body)
	err = util.BytesToStruct(bytes, &delete)

	if err = validation.Validator.Struct(delete); err != nil {
		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	results, err := database.DeleteUser(delete.IDUser)

	log.Println(results, err)
}
