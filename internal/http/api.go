package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"

	_ "github.com/marcusadriano/rinhabackend-go/docs"
)

type RinhaBackendApp struct {
	app     *fiber.App
	handler *RestHandler
	config  WebConfig
}

type WebConfig struct {
	Addr string
}

var defaultConfig = WebConfig{
	Addr: ":8080",
}

func NewRestApp(handler *RestHandler, config ...WebConfig) *RinhaBackendApp {

	cfg := defaultConfig

	if len(config) > 0 {
		cfg = config[0]
	}

	return &RinhaBackendApp{
		config:  cfg,
		handler: handler,
	}
}

//	@title			Rinha Backend API
//	@version		1.0
//	@description	Web server for "Rinha de Backend" Challenge.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8080
func (r *RinhaBackendApp) Run() error {

	app := fiber.New()
	r.app = app

	r.registerHandlers()
	r.registerSwagger()

	log.Info().Msg("Starting server on port " + r.config.Addr)

	return r.app.Listen(r.config.Addr)
}

func (r *RinhaBackendApp) registerHandlers() {
	app := r.app
	app.Post("/pessoas", r.handler.CreatePerson)
	app.Get("/pessoas/:id", r.handler.FindPersonById)
	app.Get("/pessoas", r.handler.FindAllByT)
	app.Get("/contagem-pessoas", r.handler.CountPeople)
}

func (r *RinhaBackendApp) registerSwagger() {
	app := r.app
	app.Get("/swagger/*", swagger.HandlerDefault)
}
