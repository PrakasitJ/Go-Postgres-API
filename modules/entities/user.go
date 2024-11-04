package entities

type UsersUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	GetUsers() ([]UserGetAllRes, error)
}

type UsersRepository interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	GetUsers() ([]UserGetAllRes, error)
}

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	User_id      string  `json:"user_id" db:"user_id"`
	Username     string  `json:"username" db:"username"`
	Password     string  `json:"password" db:"password"`
	Salt         string  `json:"salt" db:"salt"`
	Email        string  `json:"email" db:"email"`
	First_name   string  `json:"first_name" db:"first_name"`
	Last_name    string  `json:"last_name" db:"last_name"`
	Phone_number string  `json:"phone_number" db:"phone_number"`
	Photo_url    string  `json:"photo_url" db:"photo_url"`
	Salary       float64 `json:"salary" db:"salary"`
	Priority     string  `json:"priority" db:"priority"`
}

type UserGetAllRes struct {
	User_id      string         `json:"user_id" db:"user_id"`
	Username     string         `json:"username" db:"username"`
	Email        string         `json:"email" db:"email"`
	First_name   string         `json:"first_name" db:"first_name"`
	Last_name    string         `json:"last_name" db:"last_name"`
	Phone_number string         `json:"phone_number" db:"phone_number"`
	Photo_url    any `json:"photo_url" db:"photo_url"`
	Salary       float64        `json:"salary" db:"salary"`
	Priority     string         `json:"priority" db:"priority"`
}