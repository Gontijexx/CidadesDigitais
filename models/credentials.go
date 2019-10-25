package models

import "database/sql"

type credentials struct {
	Login sql.NullString `json: "login"	validate: "alphanum"`
	Senha sql.NullString `json: "senha"	validate: "alphanum, min=8"`
}
