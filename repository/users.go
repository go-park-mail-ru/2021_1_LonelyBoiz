package repository

import (
	"fmt"
	"server/api"

	_ "github.com/jackc/pgx/stdlib"
)

type Item struct {
	ID          uint32 `schema:"-"`
	Title       string `schema:"title,required"`
	Description string `schema:"description,required"`
	CreatedBy   uint32 `schema:"-"`
}

func (repo *RepoSqlx) Add(elem *Item) (int64, error) {
	result, err := repo.DB.NamedExec(
		`INSERT INTO person (first_name,last_name,email) VALUES (:title, :description)`,
		map[string]interface{}{
			"title":       elem.Title,
			"description": elem.Description,
		})
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *RepoSqlx) AddUser(newUser api.User) (int64, error) {
	id, err := repo.DB.NamedExec(
		`INSERT INTO users (
			email, 
			name, 
			passwordHash, 
			birthday, 
			description, 
			city,	
			sex, 
			isactive, 
			isdeleted
			) 
		VALUES (
			:email, 
			:name, 
			:pass, 
			:birth, 
			:description, 
			:city, 
			:sex, 
			:isActive, 
			:isDeleted
		)`,
		map[string]interface{}{
			"email":       newUser.Email,
			"pass":        newUser.PasswordHash,
			"name":        newUser.Name,
			"birth":       newUser.Birthday,
			"description": newUser.Description,
			"city":        newUser.City,
			"sex":         newUser.Sex,
			"isActive":    true,
			"isDeleted":   false,
		})

	if err != nil {
		return -1, err
	}

	return id.LastInsertId()
}

func (repo *RepoSqlx) GetUser(id int) (*api.User, error) {
	var user []api.User
	err := repo.DB.Select(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	fmt.Println("user=", user, "id=", id)

	return &user[0], nil
}

func (repo *RepoSqlx) DeleteUser(id int) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
		SET isDeleted = TRUE
		WHERE userid = $1`,
		id,
	)

	return err
}

func (repo *RepoSqlx) ChangeUser(newUser api.User) error {
	_, err := repo.DB.Exec(
		`UPDATE users 
			SET email = $1, name = $2, passwordHash = $3,
			birthday = $4, description = $5, city = $6,
			sex = $7, isActive = $8, isDeleted = $9`,
		newUser.Email, newUser.Name, newUser.PasswordHash,
		newUser.Birthday, newUser.Description, newUser.City,
		newUser.Sex, true, false,
	)

	return err
}
