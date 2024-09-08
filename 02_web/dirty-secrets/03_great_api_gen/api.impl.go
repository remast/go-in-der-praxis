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

func (Server) GetApiDirtySecrets(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(repository.GetAll())
}
