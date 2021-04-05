package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	model "server/internal/pkg/models"

	_ "github.com/jackc/pgx/stdlib"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (repo *UserRepository) AddUser(newUser model.User) (int, error) {
	var id int

	log.Println("bd")

	err := repo.DB.QueryRowx(
		`INSERT INTO users (
			email, 
			name,
			passwordHash,
			birthday,
			description,
			city,
			sex,
			datePreference,
			isActive,
			isDeleted,
			photos
		) VALUES (
			$1, 
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		) RETURNING id`,
		newUser.Email,
		newUser.Name,
		newUser.PasswordHash,
		newUser.Birthday,
		newUser.Description,
		newUser.City,
		newUser.Sex,
		newUser.DatePreference,
		newUser.IsActive,
		newUser.IsDeleted,
		newUser.Photos,
	).Scan(&id)
	if err != nil {
		return -1, err
	}

	log.Println(id)

	return id, nil
}

func (repo *UserRepository) GetUser(id int) (model.User, error) {
	var user []model.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return model.User{}, err
	}

	return user[0], nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
		SET isDeleted = TRUE
		WHERE id = $1`,
		id,
	)

	return err
}

func (repo *UserRepository) ChangeUser(newUser model.User) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
			SET email = $1, name = $2, passwordHash = $3,
			birthday = $4, description = $5, city = $6,
			sex = $7, datePreference = $8, isActive = $9, 
			isDeleted = $10, photos = $11
		WHERE id = $12`,
		newUser.Email, newUser.Name, newUser.PasswordHash,
		newUser.Birthday, newUser.Description, newUser.City,
		newUser.Sex, newUser.DatePreference, newUser.IsActive,
		newUser.IsDeleted, newUser.Photos, newUser.Id,
	)

	return err
}

func (repo *UserRepository) CheckMail(email string) bool {
	var user []model.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return false
	}
	if len(user) == 0 {
		return false
	}

	return true
}

func (repo *UserRepository) GetPass(id int) ([]byte, error) {
	var pass [][]byte
	err := repo.DB.Select(&pass, `SELECT passwordHash FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	return pass[0], nil
}

func (repo *UserRepository) SignIn(email string) (model.User, error) {
	var user []model.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return model.User{}, err
	}

	if len(user) == 0 {
		return model.User{}, errors.New("пользователь не найден")
	}

	return user[0], nil
}
