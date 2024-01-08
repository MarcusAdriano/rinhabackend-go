package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	pb "github.com/marcusadriano/rinhabackend-go/internal/cache"
	"github.com/marcusadriano/rinhabackend-go/internal/db/postgres"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PessoaCachedRepository struct {
	source PessoaRepository
	grpc   pb.CacheServiceClient

	grpcTimeout time.Duration
}

type CacheConfig struct {
	Addr          string
	Source        PessoaRepository
	ClientTimeout time.Duration
}

func NewPessoaCachedRepository(config CacheConfig) PessoaRepository {

	conn, err := grpc.Dial(config.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msgf("cache did not connect (will using 'direct connection'): %v", err)
		return config.Source
	}

	client := pb.NewCacheServiceClient(conn)

	return &PessoaCachedRepository{
		source:      config.Source,
		grpc:        client,
		grpcTimeout: config.ClientTimeout,
	}
}

func (p *PessoaCachedRepository) CountPeople(ctx context.Context) (int64, error) {
	return p.source.CountPeople(ctx)
}

func (p *PessoaCachedRepository) CreatePerson(ctx context.Context, person postgres.CreatePessoaParams) error {

	cache, err := p.grpc.Get(ctx, &pb.GetRequest{
		Key: person.ID.String(),
	})

	if err == nil && cache.Value != nil {
		return errors.New("person already exists")
	}

	err = p.source.CreatePerson(ctx, person)
	if err != nil {
		return err
	}

	if data, err := json.Marshal(person); err == nil {
		p.grpc.Put(ctx, &pb.PutRequest{
			Key:   person.ID.String(),
			Value: data,
		})
	}

	return nil
}

func (p *PessoaCachedRepository) FindAllByTerm(ctx context.Context, search string) ([]postgres.Pessoa, error) {
	return p.source.FindAllByTerm(ctx, search)
}

func (p *PessoaCachedRepository) FindPersonById(ctx context.Context, id string) (postgres.Pessoa, error) {

	found, err := p.grpc.Get(ctx, &pb.GetRequest{
		Key: id,
	})

	if err != nil {
		return p.source.FindPersonById(ctx, id)

	}

	var person postgres.Pessoa
	err = json.Unmarshal(found.Value, &person)
	return person, err
}
