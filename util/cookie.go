package util

import (
	"math/rand"
	"net/http"

	"github.com/gorilla/sessions"
)

//struct que ira guardar o nome do cookie
type Cookie struct {
	Name string
}

//cria uma string aleatoria que sera usada como nome do cookie
func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//atribui uma string aleatoria para a struct Cookie e retorna esse valor
func generateCookie() string {
	var aux *Cookie
	aux.Name = randStringRunes(10)
	return aux.Name
}

func Session(w http.ResponseWriter, r *http.Request) {

	key := []byte(randStringRunes(10))
	store := sessions.NewCookieStore(key)

	cookie := generateCookie()

	session, _ := store.Get(r, cookie)

	session.Values["authenticated"] = true
	session.Save(r, w)
}
