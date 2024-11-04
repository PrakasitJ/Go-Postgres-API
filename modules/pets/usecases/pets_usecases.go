package usecases

import (
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
)

type petsUse struct {
	PetsRepo entities.PetRepository
}

// Constructor
func NewPetsUsecase(petsRepo entities.PetRepository) entities.PetUsecase {
	return &petsUse{
		PetsRepo: petsRepo,
	}
}

func (u *petsUse) GetPets() ([]entities.PetGetAllRes, error) {
	// Send req next to repository
	pets, err := u.PetsRepo.GetPets()
	if err != nil {
		return nil, err
	}
	return pets, nil
}
