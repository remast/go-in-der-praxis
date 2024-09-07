package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var repository DirtySecretRepository = NewDirtySecretRepository()

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dirty Secrets!")
}

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

func main() {
	// 1. Router erzeugen
	router := http.NewServeMux()

	// 2. Handler registrieren
	router.HandleFunc("GET /api", getHandler)
	router.HandleFunc("POST /api", postHandler)
	router.HandleFunc("GET /api/{id}", getByIdHandler)
	router.HandleFunc("PUT /api/{id}", putHandler)
	router.HandleFunc("DELETE /api/{id}", deleteHandler)

	router.HandleFunc("/", helloHandler)

	// 3. Server mit Router starten
	http.ListenAndServe(":8080", router)
}
