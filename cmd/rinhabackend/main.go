package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusadriano/rinhabackend-go/internal/http"
	"github.com/marcusadriano/rinhabackend-go/internal/repository"
	"github.com/marcusadriano/rinhabackend-go/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Msgf("Error connecting to database: %v", err)
		panic(err)
	}
	defer pool.Close()

	repo := repository.NewPessoaRepository(pool)
	srv := service.NewPessoaService(repo)

	handler := http.NewRestHandler(srv)

	port := os.Getenv("SERVER_PORT")
	app := http.NewRestApp(handler, http.WebConfig{
		Addr: ":" + port,
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := app.Run(); err != nil {
		log.Error().Msgf("Error to init app: %v", err)
		os.Exit(1)
	}
}
