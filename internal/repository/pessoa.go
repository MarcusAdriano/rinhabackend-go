package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusadriano/rinhabackend-go/internal/db/postgres"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
	"strings"
)

type PessoaRepository interface {
	CreatePerson(ctx context.Context, person *model.CreatePerson) (postgres.Pessoa, error)
	FindPersonById(ctx context.Context, id string) (postgres.Pessoa, error)
	FindAllByTerm(ctx context.Context, search string) ([]postgres.Pessoa, error)
	CountPeople(ctx context.Context) (int64, error)
}

type pessoaRepository struct {
	pool *pgxpool.Pool
}

func NewPessoaRepository(pool *pgxpool.Pool) PessoaRepository {
	return pessoaRepository{
		pool: pool,
	}
}

func (p pessoaRepository) CreatePerson(ctx context.Context, person *model.CreatePerson) (postgres.Pessoa, error) {
	q := postgres.New(p.pool)

	id := uuid.New()

	params := postgres.CreatePessoaParams{
		ID:      id,
		Nome:    person.Nome,
		Apelido: person.Apelido,
		Stack:   pgtype.Text{String: strings.Join(person.Stack, ","), Valid: true},
	}

	return q.CreatePessoa(ctx, params)
}

func (p pessoaRepository) FindPersonById(ctx context.Context, id string) (postgres.Pessoa, error) {
	q := postgres.New(p.pool)

	return q.GetPessoa(ctx, uuid.MustParse(id))
}

func (p pessoaRepository) FindAllByTerm(ctx context.Context, search string) ([]postgres.Pessoa, error) {
	q := postgres.New(p.pool)

	return q.SearchPessoas(ctx, search)
}

func (p pessoaRepository) CountPeople(ctx context.Context) (int64, error) {
	q := postgres.New(p.pool)
	return q.CountPessoas(ctx)
}
