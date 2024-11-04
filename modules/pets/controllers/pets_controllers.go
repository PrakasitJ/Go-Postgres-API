package controllers

import (
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type petsController struct {
	PetsUse entities.PetUsecase
}

func NewPetsController(r fiber.Router, petsUse entities.PetUsecase) {
	controllers := &petsController{
		PetsUse: petsUse,
	}
	r.Get("/getAllPet", controllers.GetPets)
}

func (c *petsController) GetPets(ctx *fiber.Ctx) error {
	pets, err := c.PetsUse.GetPets()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(pets)
}


