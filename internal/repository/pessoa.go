package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusadriano/rinhabackend-go/internal/db/postgres"
)

type PessoaRepository interface {
	CreatePerson(ctx context.Context, person postgres.CreatePessoaParams) error
	FindPersonById(ctx context.Context, id string) (postgres.GetPessoaRow, error)
	FindAllByTerm(ctx context.Context, search string) ([]postgres.SearchPessoasRow, error)
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

func (p pessoaRepository) CreatePerson(ctx context.Context, params postgres.CreatePessoaParams) error {
	q := postgres.New(p.pool)

	return q.CreatePessoa(ctx, params)
}

func (p pessoaRepository) FindPersonById(ctx context.Context, id string) (postgres.GetPessoaRow, error) {
	q := postgres.New(p.pool)

	return q.GetPessoa(ctx, uuid.MustParse(id))
}

func (p pessoaRepository) FindAllByTerm(ctx context.Context, search string) ([]postgres.SearchPessoasRow, error) {
	q := postgres.New(p.pool)

	return q.SearchPessoas(ctx, search)
}

func (p pessoaRepository) CountPeople(ctx context.Context) (int64, error) {
	q := postgres.New(p.pool)
	return q.CountPessoas(ctx)
}
