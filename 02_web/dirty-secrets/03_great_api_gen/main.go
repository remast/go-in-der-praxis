package main

import (
	"net/http"
)

func main() {
	// 1. Router erzeugen
	router := http.NewServeMux()

	// 2. API-Server erzeugen
	apiServer := NewServer()
	apiHandler := HandlerFromMux(apiServer, router)

	// 3. HTTP-Server erzeugen
	httpServer := &http.Server{
		Handler: apiHandler,
		Addr:    "0.0.0.0:8080",
	}

	// 4. HTTP-Server starten
	httpServer.ListenAndServe()
}
