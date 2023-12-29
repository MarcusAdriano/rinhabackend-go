package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
	"github.com/marcusadriano/rinhabackend-go/internal/service"
	"github.com/rs/zerolog/log"
	"strconv"
)

type RestHandler struct {
	srv service.PessoaService
}

func NewRestHandler(srv service.PessoaService) *RestHandler {
	return &RestHandler{
		srv: srv,
	}
}

func sendError(ctx *fiber.Ctx, status int, err error) error {
	log.Error().Msgf("Error: %s", err.Error())
	return ctx.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// CreatePerson godoc
//
//	@Summary		Create a new person
//	@Description	Create a new person with the input payload
//	@Tags			pessoas
//	@Accept			json
//	@Produce		json
//	@Param			person	body		model.CreatePerson	true	"Pessoa"
//	@Success		200		{object}	model.CreatePerson
//	@Failure		400		{object}	string
//	@Router			/pessoas [post]
func (r *RestHandler) CreatePerson(c *fiber.Ctx) error {
	c.Accepts("application/json")

	request := new(model.CreatePerson)
	if err := c.BodyParser(request); err != nil {
		return sendError(c, fiber.StatusBadRequest, err)
	}

	// Validations
	if len(request.Nome) == 0 || len(request.Nome) > 100 {
		return sendError(c, fiber.StatusUnprocessableEntity, errors.New("nome nao pode ser vazio ou possuir mais de 100 caracteres"))
	}

	if len(request.Apelido) == 0 || len(request.Apelido) > 32 {
		return sendError(c, fiber.StatusUnprocessableEntity, errors.New("apelido nao pode ser vazio ou possuir mais de 32 caracteres"))
	}

	if len(request.Nascimento) != 10 {
		return sendError(c, fiber.StatusBadRequest, errors.New("nascimento deve estar no formato yyyy-mm-dd"))
	}

	p, err := r.srv.CreatePerson(c.Context(), request)
	if err != nil {
		if err.Error() == "invalid stack" {
			return sendError(c, fiber.StatusBadRequest, err)
		}
		return sendError(c, fiber.StatusUnprocessableEntity, err)
	}

	c.Set("Location", "/pessoas/"+p.ID)
	c.Status(fiber.StatusCreated)
	return nil
}

// FindPersonById godoc
//
//	@Summary		Find a person by ID
//	@Description	Get a person by ID
//	@Tags			pessoas
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Person ID"
//	@Success		200	{object}	model.PersonResponse
//	@Router			/pessoas/{id} [get]
func (r *RestHandler) FindPersonById(c *fiber.Ctx) error {

	id := c.Params("id")
	person, err := r.srv.FindPersonById(c.Context(), id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.Status(fiber.StatusNotFound)
			return nil
		}
		return sendError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(person)
}

// FindAllByT godoc
//
//	@Summary		Find all people by term
//	@Description	Get all people by term
//	@Tags			pessoas
//	@Accept			json
//	@Produce		json
//	@Param			t	query	string	true	"search term"
//	@Success		200	{array}	model.PersonResponse
//	@Router			/pessoas [get]
func (r *RestHandler) FindAllByT(c *fiber.Ctx) error {
	t := c.Query("t")

	people, err := r.srv.FindAllByTerm(c.Context(), t)
	if err != nil {
		return sendError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(people)
}

// CountPeople godoc
//
//	@Summary		Count people
//	@Description	Count people
//	@Tags			pessoas
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	number
//	@Router			/contagem-pessoas [get]
func (r *RestHandler) CountPeople(c *fiber.Ctx) error {
	count, err := r.srv.CountPeople(c.Context())
	if err != nil {
		return sendError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendString(strconv.FormatInt(count, 10))
}
