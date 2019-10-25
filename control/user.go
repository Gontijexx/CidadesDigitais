package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"log"
	"net/http"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	var user models.User

	results, err, db := database.ListaGeral("usuario", "cod_usuario")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for results.Next() {
		// for each row, scan the result into our tag composite object
		err := results.Scan(&user.IDUser, &user.Nome, &user.Email, &user.Login, &user.Status, &user.Senha)
		if err != nil {
			log.Printf("[WARN] Could not SCAN in database, because: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("[INFO] Listing users: %v %v", err, user)
		bytes, err := util.StructToBytes(user)

		w.Write(bytes)
	}

	defer db.Close()
}
