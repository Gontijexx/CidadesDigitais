package util

import "math/rand"

//struct que ira guardar o nome do cookie
type Cookie struct {
	Name string
}

//cria uma string aleatoria que sera usada como nome do cookie
func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//atribui uma string aleatoria para a struct Cookie e retorna esse valor
func GenerateCookie() string {
	var aux Cookie
	aux.Name = randStringRunes(10)
	return aux.Name
}
