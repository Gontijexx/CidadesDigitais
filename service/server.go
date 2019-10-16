package service

import (
	"CidadesDigitais/config"
	"CidadesDigitais/validation"
	"log"
	"net/http"
	"time"
)

func CreateServer() (server *http.Server) {

	server = &http.Server{
		Addr: config.SERVER_ADDR,

		IdleTimeout:  200 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	}

	return
}

func StopServer() {}

func StartServer() {
	s := CreateServer()

	h := CreateHandler()

	s.Handler = h

	validation.CreateValidator()

	log.Fatal(s.ListenAndServe())

}
