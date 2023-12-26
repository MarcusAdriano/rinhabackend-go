package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcusadriano/rinhabackend-go/internal/model"
)

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
func CreatePerson(c *fiber.Ctx) error {
	c.Accepts("application/json")

	request := new(model.CreatePerson)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("")
	}

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
func FindPersonById(c *fiber.Ctx) error {
	return c.SendString("[]")
}

// FindAllByT godoc
//
//	@Summary		Find all people by T
//	@Description	Get all people by T
//	@Tags			pessoas
//	@Accept			json
//	@Produce		json
//	@Param			t	query	string	true	"T"
//	@Success		200	{array}	model.PersonResponse
//	@Router			/pessoas [get]
func FindAllByT(c *fiber.Ctx) error {
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
func CountPeople(c *fiber.Ctx) error {
	return c.SendString("10")
}
