package model

import "time"

type CreatePerson struct {
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento time.Time `json:"nascimento"`
	Stack      []string  `json:"stack"`
}

type PersonResponse struct {
	ID         string    `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento time.Time `json:"nascimento"`
	Stack      []string  `json:"stack"`
}
