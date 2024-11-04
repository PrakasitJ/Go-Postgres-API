package entities

type PetUsecase interface {
	GetPets() ([]PetGetAllRes, error)
}

type PetRepository interface {
	GetPets() ([]PetGetAllRes, error)
}

type PetGetAllRes struct {
	Pet_id      string         `json:"pet_id" db:"pet_id"`
	Pet_name    string         `json:"pet_name" db:"pet_name"`
	Age_years   int            `json:"age_years" db:"age_years"`
	Age_months  int            `json:"age_months" db:"age_months"`
	Species     string         `json:"species" db:"species"`
	Breed       any         `json:"breed" db:"breed"`
	Sex         any         `json:"sex" db:"sex"`
	Photo_url   any         `json:"photo_url" db:"photo_url"`
	Weight      float64        `json:"weight" db:"weight"`
	Adopted     bool           `json:"adopted" db:"adopted"`
	Spayed      bool           `json:"spayed" db:"spayed"`
	Description any         `json:"description" db:"description"`
	Color       any `json:"color" db:"color"`
}