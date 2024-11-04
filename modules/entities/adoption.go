package entities


type AdoptionUsecase interface {
	GetAdoptions() ([]AdoptionGetAllRes, error)
	GetAdoptionByID(addedID int64) (AdoptionGetAllRes, error)
}

type AdoptionRepository interface {
	GetAdoptions() ([]AdoptionGetAllRes, error)
	GetAdoptionByID(addedID int64) (AdoptionGetAllRes, error)
}

type AdoptionGetByIDReq struct {
	Added_id int64 `json:"added_id" db:"added_id"`
}

type AdoptionGetAllRes struct {
	Added_id     int64  `json:"added_id" db:"added_id"`
	Added_user   any `json:"added_user" db:"added_user"`
	Pet_id       any  `json:"pet_id" db:"pet_id"`
	Added_date   any `json:"added_date" db:"added_date"`
	Adopted_date any `json:"adopted_date" db:"adopted_date"`
}