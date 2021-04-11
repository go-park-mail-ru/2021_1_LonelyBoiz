package repository

import (
	"errors"
	"fmt"
	"log"
	model "server/internal/pkg/models"

	"github.com/jmoiron/sqlx"

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
			isDeleted
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
			$10
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
	err = repo.DB.Select(&user[0].Photos, `SELECT * FROM photos WHERE userid = $1`, id)
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
			SET email = $1, name = $2, birthday = $3, 
			description = $4, city = $5, sex = $6, 
			datePreference = $7, isActive = $8,	photos = $9
		WHERE id = $10`,
		newUser.Email, newUser.Name, newUser.Birthday,
		newUser.Description, newUser.City, newUser.Sex,
		newUser.DatePreference, newUser.IsActive, newUser.Photos,
		newUser.Id,
	)

	return err
}

func (repo *UserRepository) CheckMail(email string) (bool, error) {
	var emails []string
	err := repo.DB.Select(&emails, `SELECT email FROM users WHERE email = $1`, email)
	fmt.Println(err, emails)
	if err != nil {
		return true, err
	}
	if len(emails) == 0 {
		return false, nil
	}

	return true, nil
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

func (repo *UserRepository) AddPhoto(userId int) (int, error) {
	var photoId int

	err := repo.DB.QueryRowx(
		`INSERT INTO photos (userid) VALUES (&1) RETURNING photoId;`,
		userId,
	).Scan(&photoId)
	if err != nil {
		return -1, err
	}

	return photoId, nil
}

func (repo *UserRepository) CheckPhoto(photoId int, userId int) (bool, error) {
	var idFromDB []int
	err := repo.DB.Select(&idFromDB, `SELECT * FROM photos WHERE photoId = $1`, photoId)
	if err != nil {
		return false, err
	}
	if len(idFromDB) == 0 || idFromDB[0] != userId {
		return false, nil
	}

	return true, nil
}

func (repo *UserRepository) ChangePassword(userId int, hash []byte) error {
	_, err := repo.DB.Exec(
		`UPDATE users SET passwordHash = $1 WHERE id = $2`,
		hash, userId,
	)

	return err
}
