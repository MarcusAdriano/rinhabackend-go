package model

type CreatePerson struct {
	Apelido    string `json:"apelido"`
	Nome       string `json:"nome"`
	Nascimento string `json:"nascimento"`
	Stack      *[]any `json:"stack"`
}

type PersonResponse struct {
	ID         string   `json:"id"`
	Apelido    string   `json:"apelido"`
	Nome       string   `json:"nome"`
	Nascimento string   `json:"nascimento"`
	Stack      []string `json:"stack"`
}
