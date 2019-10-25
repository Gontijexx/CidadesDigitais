//go run main.go
//no navegador abrir http://localhost:8080/secret que exibe "Forbidden"
//no navegador abrir http://localhost:8080/login para o usuario entrar
//no navegador abrir http://localhost:8080/secret como o usuario esta logado agora printa o cookie no terminal
//no navegador abrir http://localhost:8080/secret se apertar F5 gera um novo cookie
//no navegador abrir http://localhost:8080/logout para o usuario sair
//no navegador abrir http://localhost:8080/secret para de exibir o cookie no terminal
//repita o processo

package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	key         = securecookie.GenerateRandomKey(32)
	store       = sessions.NewCookieStore(key)
)

func Cookie() (aux string) {
	var Aux = RandStringRunes(10)
	return Aux
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = LetterRunes[rand.Intn(len(LetterRunes))]
	}
	return string(b)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	//	aux := RandStringRunes(10)
	session, _ := store.Get(r, Aux)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var Aux = Cookie()
	session, _ := store.Get(r, Aux)

	fmt.Print(session)

	//Autenticacao de usuario

	session.Values["authenticated"] = true
	session.Save(r, w)

}

func logout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, aux)

	//Revogar autenticacao de usuario

	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
