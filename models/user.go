package models

import "database/sql"

type User struct {
	IDUser     int          `json: "id"		validate: "required"`
	Nome       string       `json: "nome"	validate: "alphanum"`
	Email      string       `json: "email"	validate: "alphanum, email"`
	Status     sql.NullBool `json: "status"	validate: "alphanum"`
	Login      string       `json: "login"	validate: "alphanum"`
	Senha      string       `json: "senha"	validate: "alphanum, min=8"`
	ModuloUser []ModuloUser `json: "modulouser"`
}

type ModuloUser struct {
	CodMod int `json: "codmod"`
}

type Credentials struct {
	Login string `json: "login"	validate: "alphanum, required"`
	Senha string `json: "senha"	validate: "min=8, required"`
}
