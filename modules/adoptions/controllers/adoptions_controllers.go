package controllers

import (
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type adoptionsController struct {
	AdoptionsUse entities.AdoptionUsecase
}

func NewAdoptionsController(r fiber.Router, adoptionsUse entities.AdoptionUsecase) {
	controllers := &adoptionsController{
		AdoptionsUse: adoptionsUse,
	}
	r.Get("/getAllAdoption", controllers.GetAdoptions)
	r.Post("/getAdoptionByID", controllers.GetAdoptionByID)
}

func (c *adoptionsController) GetAdoptions(ctx *fiber.Ctx) error {
	adoptions, err := c.AdoptionsUse.GetAdoptions()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(adoptions)
}

func (c *adoptionsController) GetAdoptionByID(ctx *fiber.Ctx) error {
	var req entities.AdoptionGetByIDReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format",
		})
	}

	adoption, err := c.AdoptionsUse.GetAdoptionByID(req.Added_id) // Use the ID from the request
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(adoption)
}




