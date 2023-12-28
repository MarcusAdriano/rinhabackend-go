package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
	"github.com/marcusadriano/rinhabackend-go/internal/service"
)

type RestHandler struct {
	srv service.PessoaService
}

func NewRestHandler(srv service.PessoaService) *RestHandler {
	return &RestHandler{
		srv: srv,
	}
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	p, err := r.srv.CreatePerson(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Location", "/pessoas/"+p.ID)
	return c.SendString("")
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
	return c.SendString("[]")
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
	//t := c.Query("t")
	return c.SendString("[]")
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
	return c.SendString("10")
}
