package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for {
		password := getPassword()
		hash := hashAndSalt(password)

		password2 := getPassword()
		hash2 := hashAndSalt(password2)
		fmt.Println(hash2)
		passwordMatch := comparePasswords(hash, password2)

		fmt.Println(":", passwordMatch)
	}
}

func getPassword() []byte {
	var password string

	fmt.Println("Entre com uma senha:")

	_, err := fmt.Scan(&password)
	if err != nil {
		log.Println(err)
	}

	return []byte(password)
}

//criptografia de senha
//recebe como parametro 'password' em []byte
//e retorna uma string criptografada
func hashAndSalt(password []byte) string {

	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(hash)

	return string(hash)
}

//compara senhas
//os parametros de compare são string e []byte
//byteHash converte o parametro string em []byte
//compareHashAndPassword compara dois []byte e retorna um bool
func comparePasswords(hashedPassword string, plainPassword []byte) bool {

	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
