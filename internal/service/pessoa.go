package service

import (
	"context"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
	"github.com/marcusadriano/rinhabackend-go/internal/repository"
	"time"
)

type PessoaService interface {
	CreatePerson(ctx context.Context, person *model.CreatePerson) (*model.PersonResponse, error)
	FindPersonById(ctx context.Context, id string) (*model.PersonResponse, error)
	FindAllByTerm(ctx context.Context, search string) ([]*model.PersonResponse, error)
	CountPeople(ctx context.Context) (int64, error)
}

type pessoaService struct {
	repo repository.PessoaRepository
}

func (p pessoaService) CreatePerson(ctx context.Context, req *model.CreatePerson) (*model.PersonResponse, error) {
	person, err := p.repo.CreatePerson(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.PersonResponse{
		ID:         person.ID.String(),
		Apelido:    "",
		Nome:       "",
		Nascimento: time.Time{},
		Stack:      nil,
	}, nil
}

func (p pessoaService) FindPersonById(ctx context.Context, id string) (*model.PersonResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p pessoaService) FindAllByTerm(ctx context.Context, search string) ([]*model.PersonResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p pessoaService) CountPeople(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewPessoaService(repo repository.PessoaRepository) PessoaService {
	return &pessoaService{
		repo: repo,
	}
}
