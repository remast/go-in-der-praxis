package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "crossnative.com/dirty-secrets/docs"
	"github.com/go-chi/chi/v5"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
)

var repository DirtySecretRepository = NewDirtySecretRepository()

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dirty Secrets!")
}

// ListAccounts lists all existing accounts
//
//	@Summary		List dirty secrets
//	@Description	Get's all known dirty secrets
//	@Tags			dirty-secrets
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		DirtySecret
//	@Failure		404	{object}	string
//	@Router			/api/dirty-secrets [get]
func getHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Fehler speichern
	err := json.NewEncoder(w).Encode(repository.GetAll())

	// 2. Fehler prüfen und ggf. behandeln
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getByIdHandler(w http.ResponseWriter, req *http.Request) {
	// 1. Id aus Pfad lesen
	id := req.PathValue("id")

	// 2. Prüfen ob Secret vorhanden
	if repository.ExistsById(id) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Secret not found.")
		return
	}

	// 3. Secret laden und zurück geben
	json.NewEncoder(w).Encode(repository.GetById(id))
}

func putHandler(w http.ResponseWriter, req *http.Request) {
	// 1. Id aus Pfad lesen
	id := req.PathValue("id")

	// 2. Secret aus dem HTTP-Body lesen
	var secret DirtySecret
	json.NewDecoder(req.Body).Decode(&secret)

	// 3. Secret aktualsieren
	updatedSecret := repository.Update(id, secret)

	// 4. Aktualisiertes Secret zurückgeben
	json.NewEncoder(w).Encode(updatedSecret)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Variable für Secret initialisieren
	var secret DirtySecret

	// 2. Secret aus dem HTTP-Body lesen
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Neues Secret hinzufügen
	savedSecret := repository.Save(secret)

	// 4. Gespeichertes Secret zurückgeben
	json.NewEncoder(w).Encode(savedSecret)
}

//	@title			Dirty Secrets API
//	@version		1.0
//	@description	Keeps track of dirty secrets

//	@contact.name	Jan Stamer
//	@contact.url	https://www.crossnative.com

// @tag.name dirty-secrets
// @tag.description Dirty Secrets

// @host		localhost:8080
// @BasePath	/api
func main() {
	// 1. Router erzeugen
	router := chi.NewRouter()

	// 2. Handler registrieren
	router.HandleFunc("GET /api/dirty-secrets", getHandler)
	router.HandleFunc("POST /api/dirty-secrets", postHandler)
	router.HandleFunc("GET /api/dirty-secrets/{id}", getByIdHandler)
	router.HandleFunc("PUT /api/dirty-secrets/{id}", putHandler)
	router.HandleFunc("DELETE /api/dirty-secrets/{id}", deleteHandler)

	// UI mit API Dokumentation

	doc := doc.Doc{
		Title:       "Dirty Secrets API",
		Description: "Dirts Secrets API Description",
		SpecFile:    "./docs/swagger.yaml",
		SpecPath:    "/openapi/openapi.yaml",
		DocsPath:    "/openapi/docs",
		Theme:       "dark",
	}
	router.Handle("GET /openapi/*", doc.Handler())

	router.HandleFunc("/", helloHandler)

	// 3. Server mit Router starten
	http.ListenAndServe(":8080", router)
}
