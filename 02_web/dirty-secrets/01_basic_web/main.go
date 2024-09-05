package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Template kompilieren
	tpl, _ := template.New("html").Parse(`
		<html>
		  <body>
		    <h1>Dirty Secrets</h1>
			<dl>
			<!-- Schleife über secrets -->
			{{ range . }}
			  <dt>{{ .Name }}</dt>
			  <dd>{{ .Secret }}</dd>
			{{ end }}
			</dl>
		  </body>
		</html>	
	`)

	// 2. Template rendern
	tpl.Execute(w, secrets)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dirty Secrets!")
}

type DirtySecret struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

var secrets = make([]DirtySecret, 0)

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Fehler speichern
	err := json.NewEncoder(w).Encode(secrets)

	// 2. Fehler prüfen und ggf. behandeln
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postHandler2(w http.ResponseWriter, r *http.Request) {
	var secret DirtySecret
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	secrets = append(secrets, secret)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Variable für Secret initialisieren
	var secret DirtySecret

	// 2. Secret aus dem HTTP-Body lesen
	json.NewDecoder(r.Body).Decode(&secret)

	// 3. Neues Secret hinzufügen
	secrets = append(secrets, secret)
}

func main() {
	// 1. Router erzeugen
	router := http.NewServeMux()

	// 2. Handler registrieren
	router.HandleFunc("/", htmlHandler)
	router.HandleFunc("GET /api", getAllHandler)
	router.HandleFunc("POST /api", postHandler)

	// 3. Server mit Router starten
	http.ListenAndServe(":8080", router)
}
