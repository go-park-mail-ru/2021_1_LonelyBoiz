package repository

import (
	"database/sql"
	"server/api"

	_ "github.com/jackc/pgx/stdlib"
)

func (repo *RepoSqlx) AddUser(newUser api.User) (int, error) {
	var id int

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

	return id, nil
}

func (repo *RepoSqlx) GetUser(id int) (api.User, error) {
	var user []api.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return api.User{}, err
	}

	user[0].PasswordHash = nil
	return user[0], nil
}

func (repo *RepoSqlx) DeleteUser(id int) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
		SET isDeleted = TRUE
		WHERE id = $1`,
		id,
	)

	return err
}

func (repo *RepoSqlx) ChangeUser(newUser api.User) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
			SET email = $1, name = $2, passwordHash = $3,
			birthday = $4, description = $5, city = $6,
			sex = $7, datePreference = $8, isActive = $9, 
			photos = $10
		WHERE id = $11`,
		newUser.Email, newUser.Name, newUser.PasswordHash,
		newUser.Birthday, newUser.Description, newUser.City,
		newUser.Sex, newUser.DatePreference, newUser.IsActive,
		newUser.Photos,
		newUser.Id,
	)

	return err
}

func (repo *RepoSqlx) CheckMail(email string) (bool, error) {
	var user []api.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE email = $1`, email)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func (repo *RepoSqlx) GetPass(id int) ([]byte, error) {
	var pass [][]byte
	err := repo.DB.Select(&pass, `SELECT passwordHash FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	return pass[0], nil
}

func (repo *RepoSqlx) SignIn(email string) (api.User, error) {
	var user []api.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return api.User{}, err
	}

	return user[0], nil
}
