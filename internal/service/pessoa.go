package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/marcusadriano/rinhabackend-go/internal/db/postgres"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
	"github.com/marcusadriano/rinhabackend-go/internal/repository"
	"strings"
	"time"
)

const dtFormat = "2006-01-02"

type PessoaService interface {
	CreatePerson(ctx context.Context, person *model.CreatePerson) (*model.PersonResponse, error)
	FindPersonById(ctx context.Context, id string) (*model.PersonResponse, error)
	FindAllByTerm(ctx context.Context, search string) ([]model.PersonResponse, error)
	CountPeople(ctx context.Context) (int64, error)
}

type pessoaService struct {
	repo repository.PessoaRepository
}

func (p pessoaService) CreatePerson(ctx context.Context, req *model.CreatePerson) (*model.PersonResponse, error) {

	nascimentoDt, err := time.Parse(dtFormat, req.Nascimento)
	if err != nil {
		return nil, err
	}

	stack, err := parseStack(req)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	var stackText pgtype.Text
	if stack != nil && len(stack) > 0 {
		err = stackText.Scan(strings.Join(stack, ","))
		if err != nil {
			return nil, err
		}
	}

	var nascimentoSqlDt pgtype.Date
	err = nascimentoSqlDt.Scan(nascimentoDt)
	if err != nil {
		return nil, err
	}

	params := postgres.CreatePessoaParams{
		ID:         id,
		Nome:       req.Nome,
		Apelido:    req.Apelido,
		Stack:      stackText,
		Nascimento: nascimentoSqlDt,
	}

	person, err := p.repo.CreatePerson(ctx, params)
	if err != nil {
		return nil, err
	}

	return &model.PersonResponse{
		ID:         person.ID.String(),
		Apelido:    person.Apelido,
		Nome:       person.Nome,
		Nascimento: nascimentoDt.Format(dtFormat),
		Stack:      stack,
	}, nil
}

func parseStack(req *model.CreatePerson) ([]string, error) {
	var stack []string
	if req.Stack != nil {
		if len(*req.Stack) > 0 {
			for _, s := range *req.Stack {
				if v, ok := s.(string); ok && len(v) <= 32 {
					stack = append(stack, v)
				} else {
					return nil, errors.New("invalid stack")
				}
			}
		}
	}
	return stack, nil
}

func (p pessoaService) FindPersonById(ctx context.Context, id string) (*model.PersonResponse, error) {
	person, err := p.repo.FindPersonById(ctx, id)

	var stack []string
	if person.Stack.String != "" {
		stack = strings.Split(person.Stack.String, ",")
	}

	return &model.PersonResponse{
		ID:         person.ID.String(),
		Apelido:    person.Apelido,
		Nome:       person.Nome,
		Nascimento: person.Nascimento.Time.Format(dtFormat),
		Stack:      stack,
	}, err
}

func (p pessoaService) FindAllByTerm(ctx context.Context, search string) ([]model.PersonResponse, error) {
	people, err := p.repo.FindAllByTerm(ctx, search)
	if err != nil {
		return nil, err
	}

	if people == nil {
		return []model.PersonResponse{}, nil
	}

	var peopleResponse []model.PersonResponse
	for _, person := range people {
		peopleResponse = append(peopleResponse, model.PersonResponse{
			ID:         person.ID.String(),
			Apelido:    person.Apelido,
			Nome:       person.Nome,
			Nascimento: person.Nascimento.Time.Format(dtFormat),
			Stack:      strings.Split(person.Stack.String, ","),
		})
	}

	return peopleResponse, nil
}

func (p pessoaService) CountPeople(ctx context.Context) (int64, error) {
	return p.repo.CountPeople(ctx)
}

func NewPessoaService(repo repository.PessoaRepository) PessoaService {
	return &pessoaService{
		repo: repo,
	}
}
