package main

import (
	"encoding/json"
	"net/http"
)

var repository DirtySecretRepository = NewDirtySecretRepository()

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetApiDirtySecrets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(repository.GetAll())
}
