package repositories

import (
	"fmt"
	"github.com/PrakasitJ/Go-Postgres-API.git/modules/entities"
	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
		"username",
		"password"
	)
	VALUES ($1, $2)
	RETURNING "id", "username";
	`

	// Initail a user object
	user := new(entities.UsersRegisterRes)

	// Query part
	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	// defer r.Db.Close()

	return user, nil
}

func (r *usersRepo) GetUsers() ([]entities.UserGetAllRes, error) {
	query := `
	SELECT "user_id", "username","email", "first_name", "last_name", "phone_number", "photo_url", "salary", "priority"
	FROM "user";
	`

	// Initial a users object
	users := []entities.UserGetAllRes{}

	// Query part
	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var user entities.UserGetAllRes
		if err := rows.StructScan(&user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	// defer r.Db.Close()

	return users, nil
}
