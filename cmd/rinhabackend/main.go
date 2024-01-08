package main

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusadriano/rinhabackend-go/internal/http"
	"github.com/marcusadriano/rinhabackend-go/internal/repository"
	"github.com/marcusadriano/rinhabackend-go/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Msgf("Error connecting to database: %v", err)
		panic(err)
	}

	config.ConnConfig.Config.ConnectTimeout = time.Second * 2

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Error().Msgf("Error connecting to database: %v", err)
		panic(err)
	}
	defer pool.Close()

	repo := repository.NewPessoaRepository(pool)

	cacheConfig := repository.CacheConfig{
		Addr:          os.Getenv("CACHE_ADDR"),
		Source:        repo,
		ClientTimeout: time.Millisecond * 500,
	}

	cacheRepo := repository.NewPessoaCachedRepository(cacheConfig)
	srv := service.NewPessoaService(cacheRepo)

	handler := http.NewRestHandler(srv)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	app := http.NewRestApp(handler, http.WebConfig{
		Addr: ":" + port,
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := app.Run(); err != nil {
		log.Error().Msgf("Error to init app: %v", err)
		os.Exit(1)
	}
}
