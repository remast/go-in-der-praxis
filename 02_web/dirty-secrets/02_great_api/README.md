# Grandiose REST APIs mit Go

http://localhost:8080/openapi/docs

listing 1
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dirty Secrets!")
}

func main() {
	// 1. Router erzeugen
	router := http.NewServeMux()

	// 2. Handler registrieren
	router.HandleFunc("GET /", helloHandler)

	// 3. Server mit Router starten
	http.ListenAndServe(":8080", router)
}
```


listing 2
```go
func main() {
	// 1. Router erzeugen
	router := http.NewServeMux()

	// 2. REST Handler registrieren
	router.HandleFunc("GET /api", getHandler)
	router.HandleFunc("POST /api", postHandler)
	router.HandleFunc("PUT /api/{id}", putHandler)
	router.HandleFunc("DELETE /api/{id}", deleteHandler)

	// 3. Server mit Router starten
	http.ListenAndServe(":8080", router)
}
```


listing 3
```go
var repository DirtySecretRepository = NewDirtySecretRepository()

func getHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Secrets von Repository abfragen
	secrets := repository.GetAll();

	// 2. Secrets als JSON serialisieren
	json.NewEncoder(w).Encode(secrets)
}
```

listing 4
```go
func putHandler(w http.ResponseWriter, req *http.Request) {
	// 1. Id aus Pfad lesen
	id := req.PathValue("id")

	// 2. Secret aus dem HTTP-Body lesen
	var secret DirtySecret
	json.NewDecoder(req.Body).Decode(&secret)

	// 3. Secret aktualisieren
	updatedSecret := repository.Update(id, secret)

	// 4. Aktualisiertes Secret zurückgeben
	json.NewEncoder(w).Encode(updatedSecret)
}
```

listing 5
```go
if repository.ExistsById(id) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Secret not found.")
	return
}
```
