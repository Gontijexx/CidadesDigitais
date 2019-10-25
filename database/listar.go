package database

import (
	"database/sql"
	"log"
)

//ListaGeral lista qualquer elemento do BD
func ListaGeral(param string, nID string) (results *sql.Rows, err error, db *sql.DB) {
	db = ClientSQL()

	results, err = db.Query("SELECT * FROM " + param + " ORDER BY " + nID + " ASC")
	if err != nil {
		log.Printf("[WARN] Could not 'SELECT * FROM usuario' in database, because: %v\n", err)
		return
	}

	return results, err, db
}
