package server

import (
	"net/http"
	"time"
)

func createServer() (server *http.Server) {

	server = &http.Server{
		Addr: configs.SERVER_ADDR,

		IdleTimeout:  200 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	}

	return
}

func stopServer() {}

func startSever() {
	s := createServer()

	h := createHandler()

	s.Handler = h

}
