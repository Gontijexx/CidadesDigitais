package control

import (
	"CidadesDigitais/database"
	"CidadesDigitais/models"
	"CidadesDigitais/util"
	"CidadesDigitais/validation"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

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

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	/*
		aux := randStringRunes(10)
		session, err := store.Get(r, aux)

		if err != nil {
			log.Printf("[ERROR] It was not possible to create sessions because: %v\n", err)
		}
	*/
	body := r.Body

	bytes, err := util.BodyToBytes(body)

	err = util.BytesToStruct(bytes, login)

	// checks if struct is a valid one
	if err := validation.Validator.Struct(login); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	_, err = database.CheckLogin(login.Login)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}

}

func Senha(w http.ResponseWriter, r *http.Request) {
	var password models.Senha
	var user models.User

	body := r.Body

	bytes, err := util.BodyToBytes(body)

	err = util.BytesToStruct(bytes, password)

	if err := validation.Validator.Struct(password); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	results, err := database.CheckSenha(password.Senha)

	fmt.Print(results)

	err = results.Scan(&user.IDUser, &user.Nome, &user.Email, &user.Login, &user.Status, &user.Senha)
	if err != nil {
		log.Printf("[WARN] Could not SCAN in database, because: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pass, err := util.StructToBytes(password)
	w.Write(pass)

	//verificar se os parametros estao corretos!!!
	//compara a senha que veio do banco, tranformando ela em []byte(password.Senha)
	//com a senha que vem do front
	err = bcrypt.CompareHashAndPassword([]byte(user.Senha.String), []byte(password.Senha))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}

}

func hashAndSalt(password []byte) string {

	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(hash)

	return string(hash)
}
