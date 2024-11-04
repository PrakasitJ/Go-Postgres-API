package usecases

import (
	"fmt"

	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUse) GetUsers() ([]entities.UserGetAllRes, error) {
	// Send req next to repository
	users, err := u.UsersRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

