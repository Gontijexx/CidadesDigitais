package models

import "database/sql"

type User struct {
	IDUser     int            `json: "id"		validate: "required"`
	Nome       sql.NullString `json: "nome"	validate: "alphanum"`
	Email      sql.NullString `json: "email"	validate: "alphanum, email"`
	Status     sql.NullBool   `json: "status"	validate: "alphanum"`
	Login      sql.NullString `json: "login"	validate: "alphanum"`
	Senha      sql.NullString `json: "senha"	validate: "alphanum, min=8"`
	ModuloUser []ModuloUser   `json: "modulouser"`
}

type ModuloUser struct {
	CodMod int `json: "codmod"`
}

type Login struct {
	Login string `json: "login"	validate: "alphanum, required"`
}

type Senha struct {
	Senha string `json: "senha"	validate: "min=8, required"`
}