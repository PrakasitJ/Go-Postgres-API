package repositories

import (
	"fmt"
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"github.com/jmoiron/sqlx"
)

type petsRepo struct {
	Db *sqlx.DB
}

func NewPetsRepository(db *sqlx.DB) entities.PetRepository {
	return &petsRepo{
		Db: db,
	}
}

func (r *petsRepo) GetPets() ([]entities.PetGetAllRes, error) {
	query := `
	SELECT "pet_id", "pet_name", "age_years", "age_months", "species","breed", "sex", "photo_url", "weight", "adopted", "spayed", "description", "color"
	FROM "pet";
	`
	
	pets := []entities.PetGetAllRes{}

	// Query part
	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var pet entities.PetGetAllRes
		if err := rows.StructScan(&pet); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		pets = append(pets, pet)
	}
	// defer r.Db.Close()

	return pets, nil
}
