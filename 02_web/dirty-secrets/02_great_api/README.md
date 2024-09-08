# Grandiose REST APIs mit Go

go install github.com/swaggo/swag/cmd/swag@latest   

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
if !repository.ExistsById(id) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Secret not found.")
	return
}
```


listing 6
```go
//	@title			Dirty Secrets API
//	@version		1.0
//	@description	Keeps track of dirty secrets

//	@contact.name	Jan Stamer
//	@contact.url	https://www.crossnative.com
func main() { ... }
```

listing 7
```go
//	@Summary		List dirty secrets
//	@Description	Get's all known dirty secrets
//	@Tags			dirty-secrets
//	@Produce		json
//	@Success		200	{array}	DirtySecret
//	@Router			/api/dirty-secrets [get]
func getHandler(w http.ResponseWriter, r *http.Request) {}
```


go get github.com/openapi-ui/go-openapi-ui

listing 8
```go
doc := doc.Doc{
	Title:       "Dirty Secrets API",
	Description: "Dirts Secrets API Description",
	SpecFile:    "./docs/swagger.yaml",
	SpecPath:    "/openapi/swagger.yaml",
	DocsPath:    "/openapi/docs",
}
docHandler := doc.Handler()
router.Handle("GET /openapi/docs", docHandler)
router.Handle("GET /openapi/swagger.yaml", docHandler)
```

listing 9
```go
type DirtySecret struct {
	ID     string `json:"id" minLength:"4" maxLength:"10" example:"id-123"`
	Name   string `json:"name" validate:"required" example:"Frank"`
	Secret string `json:"secret" validate:"required" example:"Has alcohol problems."`
}
```

listing 10
```yaml
basePath: /api
definitions:
  main.DirtySecret:
    properties:
      id:
        example: id-123
        maxLength: 10
        minLength: 4
        type: string
      name:
        example: Frank
        type: string
```

listing 11
```yaml
package: main
output: api.gen.go
generate:
  models: true
  std-http-server: true
```

listing 12
```go
type ServerInterface interface {
	// List dirty secrets
	// (GET /api/dirty-secrets)
	GetApiDirtySecrets(w http.ResponseWriter, r *http.Request)
}
```

listing 13
```go
type Server struct{}

func (Server) GetApiDirtySecrets(w http.ResponseWriter, r *http.Request) {
	secrets := repository.GetAll()
	json.NewEncoder(w).Encode(secrets)
}
```

listing 14
```go
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
```

listing 15
```go
if !repository.ExistsById(id) {
	problem.New(
		problem.Status(404),
		problem.Detail("Secret not found."),
		problem.Custom("id", id),
	).WriteTo(w)
	return
}
```

go get -u schneider.vip/problem 

listing 16
```go
func loggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%v requested URL %v", r.Host, r.URL)
    next.ServeHTTP(w, r)
  })
}
```

go get github.com/justinas/alice

listing 17
```go
middlewares := alice.New(
	middleware.Logger,
	middleware.Recoverer,
	middleware.Timeout(30*time.Second),
)
http.ListenAndServe(":8080", middlewares.Then(router))
```