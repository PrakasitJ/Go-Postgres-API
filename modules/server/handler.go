package servers

import (
	_usersHttp "github.com/PrakasitJ/Go-Postgres-API.git/modules/users/controllers"
	_usersRepository "github.com/PrakasitJ/Go-Postgres-API.git/modules/users/repositories"
	_usersUsecase "github.com/PrakasitJ/Go-Postgres-API.git/modules/users/usecases"

	_petsHttp "github.com/PrakasitJ/Go-Postgres-API.git/modules/pets/controllers"
	_petsRepository "github.com/PrakasitJ/Go-Postgres-API.git/modules/pets/repositories"
	_petsUsecase "github.com/PrakasitJ/Go-Postgres-API.git/modules/pets/usecases"

	_adoptionsHttp "github.com/PrakasitJ/Go-Postgres-API.git/modules/adoptions/controllers"
	_adoptionsRepository "github.com/PrakasitJ/Go-Postgres-API.git/modules/adoptions/repositories"
	_adoptionsUsecase "github.com/PrakasitJ/Go-Postgres-API.git/modules/adoptions/usecases"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/")

	//* Users group
	usersGroup := v1.Group("/user")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	petsGroup := v1.Group("/pet")
	petsRepository := _petsRepository.NewPetsRepository(s.Db)
	petsUsecase := _petsUsecase.NewPetsUsecase(petsRepository)
	_petsHttp.NewPetsController(petsGroup, petsUsecase)

	adoptionsGroup := v1.Group("/adoption")
	adoptionsRepository := _adoptionsRepository.NewAdoptionsRepository(s.Db)
	adoptionsUsecase := _adoptionsUsecase.NewAdoptionsUsecase(adoptionsRepository)
	_adoptionsHttp.NewAdoptionsController(adoptionsGroup, adoptionsUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}