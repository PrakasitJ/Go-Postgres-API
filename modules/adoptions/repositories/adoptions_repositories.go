package repositories

import (
	"fmt"
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"github.com/jmoiron/sqlx"
)

type adoptionsRepo struct {
	Db *sqlx.DB
}

func NewAdoptionsRepository(db *sqlx.DB) entities.AdoptionRepository {
	return &adoptionsRepo{
		Db: db,
	}
}

func (r *adoptionsRepo) GetAdoptions() ([]entities.AdoptionGetAllRes, error) {
	query := `
	SELECT "added_id", "added_user", "pet_id", "added_id", "adopted_date"
	FROM "adoption";
	`
	
	adoptions := []entities.AdoptionGetAllRes{}

	// Query part
	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var adoption entities.AdoptionGetAllRes
		if err := rows.StructScan(&adoption); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		adoptions = append(adoptions, adoption)
	}
	// defer r.Db.Close()

	return adoptions, nil
}

func (r *adoptionsRepo) GetAdoptionByID(addedID int64) (entities.AdoptionGetAllRes, error) {
	query := `
	SELECT "added_id", "added_user", "pet_id", "adopted_date"
	FROM "adoption"
	WHERE "added_id" = $1;
	`

	adoption := entities.AdoptionGetAllRes{}

	// Query part
	row := r.Db.QueryRowx(query, addedID)
	if err := row.StructScan(&adoption); err != nil {
		fmt.Println(err.Error())
		return entities.AdoptionGetAllRes{}, err // Return empty struct in case of error
	}

	return adoption, nil
}