package usecases

import (
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
)

type adoptionsUse struct {
	AdoptionsRepo entities.AdoptionRepository
}

// Constructor
func NewAdoptionsUsecase(adoptionsRepo entities.AdoptionRepository) entities.AdoptionUsecase {
	return &adoptionsUse{
		AdoptionsRepo: adoptionsRepo,
	}
}

func (u *adoptionsUse) GetAdoptions() ([]entities.AdoptionGetAllRes, error) {
	// Send req next to repository
	adoptions, err := u.AdoptionsRepo.GetAdoptions()
	if err != nil {
		return nil, err
	}
	return adoptions, nil
}

func (u *adoptionsUse) GetAdoptionByID(addedID int64) (entities.AdoptionGetAllRes, error) {
	// Send req next to repository
	adoption, err := u.AdoptionsRepo.GetAdoptionByID(addedID)
	if err != nil {
		return entities.AdoptionGetAllRes{}, err // Return empty struct in case of error
	}
	return adoption, nil
}


