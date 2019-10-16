package main

import (
	"CidadesDigitais/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// assure server closing at end of execution
	defer service.StopServer()

	// call start server function
	service.StartServer()

}
