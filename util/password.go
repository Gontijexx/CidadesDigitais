package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// transforma a senha string em bytes
func PasswordStringToByte(password string) []byte {
	return []byte(password)
}

// criptografa a senha em byte e rotorna uma string criptografada
func PasswordByteToHashString(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		// Write a error menssage
		log.Println(err)
	}

	return string(hash)
}

// compara a senha fornecida pelo usuario (byte) com a senha associada
// ao login armazenada no banco de dados que esta criptografada (string)
func ComparePasswords(dbPassword string, userPassword []byte) bool {

	byteDbPassword := []byte(dbPassword)

	err := bcrypt.CompareHashAndPassword(byteDbPassword, userPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
