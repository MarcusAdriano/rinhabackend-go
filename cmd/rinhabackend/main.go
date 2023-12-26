package main

import (
	"github.com/marcusadriano/rinhabackend-go/internal/http"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	app := http.NewAppWithDefaultConfig()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := app.Run(); err != nil {
		log.Error().Msgf("Error to init app: %v", err)
		os.Exit(1)
	}
}
