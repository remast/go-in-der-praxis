package main

import (
	"fmt"
	"math/rand/v2"
)

type DirtySecret struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type DirtySecretRepository struct {
	secrets []*DirtySecret
}

func NewDirtySecretRepository() DirtySecretRepository {
	return DirtySecretRepository{
		secrets: make([]*DirtySecret, 0),
	}
}

func (r DirtySecretRepository) GetAll() []*DirtySecret {
	return r.secrets
}

func (r DirtySecretRepository) ExistsById(id string) bool {
	for _, s := range r.secrets {
		if s.ID == id {
			return true
		}
	}
	return false
}

func (r DirtySecretRepository) GetById(id string) *DirtySecret {
	for _, s := range r.secrets {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func (r *DirtySecretRepository) Save(secret DirtySecret) DirtySecret {
	secret.ID = fmt.Sprintf("id-%v", rand.IntN(1000))
	r.secrets = append(r.secrets, &secret)
	return secret
}

func (r *DirtySecretRepository) Update(id string, secret DirtySecret) *DirtySecret {
	for _, s := range r.secrets {
		if s.ID == id {
			s.Name = secret.Name
			s.Secret = secret.Secret
			return s
		}
	}
	return &secret
}

func (r *DirtySecretRepository) Delete(id string) {
}
