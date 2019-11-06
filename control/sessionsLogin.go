package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"log"
	"net/http"
)

/*
var (
	LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	key         = securecookie.GenerateRandomKey(32)
	store       = sessions.NewCookieStore(key)
)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = LetterRunes[rand.Intn(len(LetterRunes))]
	}
	return string(b)
}
*/

func Login(w http.ResponseWriter, r *http.Request) {
	// tratamento dos dados vindos do front-end
	var login models.Login
	body := r.Body
	bytes, err := util.BodyToBytes(body)
	err = util.BytesToStruct(bytes, &login)

	// checks if struct is a valid one
	if err = validation.Validator.Struct(login); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed) // Status 412
		return
	}

	_, err = database.CheckLogin(login.Login)

	if err != nil {
		w.WriteHeader(http.StatusForbidden) // Status 403
	} else {
		w.WriteHeader(http.StatusAccepted) // Status 202
	}
}

func Senha(w http.ResponseWriter, r *http.Request) {
	//var password models.Senha
	var user models.Login

	log.Print(user)
	/*
		body := r.Body

		bytes, err := util.BodyToBytes(body)

		err = util.BytesToStruct(bytes, &password)

		if err := validation.Validator.Struct(password); err != nil {

			log.Printf("[WARN] invalid user information, because, %v\n", err)
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}
	*/
	/*
		results, err := database.CheckSenha(password.Senha)

		fmt.Print(results)
	*/
	/*
		err = results.Scan(&user.IDUser, &user.Nome, &user.Email, &user.Login, &user.Status, &user.Senha)
		if err != nil {
			log.Printf("[WARN] Could not SCAN in database, because: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	*/
	//	pass, err := util.StructToBytes(password)
	//	w.Write(pass)

	//compara a senha que veio do banco, tranformando ela em []byte(password.Senha)
	//com a senha que vem do front-end
	//err = bcrypt.CompareHashAndPassword([]byte(user.Senha.String), []byte(password.Senha))
	/*
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusAccepted)
			//		util.Session(w, r)
		}
	*/
}
