package database

import (
	"database/sql"
	"log"
)

func ListaUsuario() (results *sql.Rows, err error) {
	db := ClientSQL()

	results, err = db.Query("SELECT * FROM usuario ORDER BY cod_usuario ASC")
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT * FROM usuario' in database, because: %v\n", err)
		return
	}

	defer db.Close()

	return results, err
}
