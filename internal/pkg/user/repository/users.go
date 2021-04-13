package repository

import (
	"database/sql"
	"errors"
	"fmt"
	model "server/internal/pkg/models"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (repo *UserRepository) AddUser(newUser model.User) (int, error) {
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
			instagram
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
		newUser.Instagram,
	).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repo *UserRepository) GetUser(id int) (model.User, error) {
	var user []model.User
	err := repo.DB.Select(&user,
		`SELECT id, 
			email,
    		name,
    		birthday,
			instagram,
    		description,
    		city,
    		sex,
    		datepreference,
    		isactive,
    		isdeleted
		FROM users WHERE id = $1`,
		id)
	if err != nil {
		return model.User{}, err
	}
	if len(user) == 0 {
		return model.User{}, sql.ErrNoRows
	}

	err = repo.DB.Select(&user[0].Photos, `SELECT photoId FROM photos WHERE userid = $1`, id)
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}
	if len(user[0].Photos) == 0 {
		user[0].Photos = make([]int, 0)
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
			datePreference = $7, isActive = $8, instagram = $9
		WHERE id = $10`,
		newUser.Email, newUser.Name, newUser.Birthday,
		newUser.Description, newUser.City, newUser.Sex,
		newUser.DatePreference, newUser.IsActive, newUser.Instagram,
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

func (repo *UserRepository) GetPassWithEmail(email string) ([]byte, error) {
	var pass [][]byte
	err := repo.DB.Select(&pass, `SELECT passwordHash FROM users WHERE email = $1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if len(pass) == 0 {
		return nil, nil
	}

	return pass[0], nil
}

func (repo *UserRepository) GetPassWithId(id int) ([]byte, error) {
	var pass [][]byte
	err := repo.DB.Select(&pass, `SELECT passwordHash FROM users WHERE id= $1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if len(pass) == 0 {
		return nil, nil
	}

	return pass[0], nil
}

func (repo *UserRepository) SignIn(email string) (model.User, error) {
	var user []model.User
	err := repo.DB.Select(&user,
		`SELECT id,
			email,
    		name,
    		birthday,
    		description,
    		city,
    		sex,
			instagram,
			passwordhash,
    		datepreference,
    		isactive,
    		isdeleted
		FROM users WHERE email = $1`, email)
	if err != nil {
		return model.User{}, err
	}
	if len(user) == 0 {
		return model.User{}, errors.New("пользователь не найден")
	}

	return user[0], nil
}

func (repo *UserRepository) AddPhoto(userId int, image string) (int, error) {
	var photoId int

	err := repo.DB.QueryRowx(
		`INSERT INTO photos (userid, value) VALUES ($1, $2) RETURNING photoId;`,
		userId, image,
	).Scan(&photoId)
	if err != nil {
		return -1, err
	}

	return photoId, nil
}

func (repo *UserRepository) GetPhoto(userId int, photoId int) (string, error) {
	var image []string

	err := repo.DB.Select(&image,
		`SELECT value FROM photos WHERE userid = $1 and photoId = $2`,
		userId, photoId,
	)
	if err != nil {
		return "", err
	}

	return image[0], nil
}

func (repo *UserRepository) CheckPhoto(photoId int, userId int) (bool, error) {
	var idFromDB []int
	err := repo.DB.Select(&idFromDB, `SELECT userId FROM photos WHERE photoId = $1`, photoId)
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

func (repo *UserRepository) CreateChat(userId1 int, userId2 int) (int, error) {
	var chatId int
	err := repo.DB.QueryRow(
		`INSERT INTO chats (userid1, userid2)
			VALUES (
        	$1,
        	$2
    	) RETURNING chatid`,
		userId1, userId2,
	).Scan(&chatId)
	if err != nil {
		return -1, err
	}

	return chatId, err
}

func (repo *UserRepository) GetChatById(chatId int, userId int) (model.Chat, error) {
	var chats []model.Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id AS chatId,
    		users.id AS partnerId,
    		users.name AS partnerName,
    		lastMessage.text AS lastMessage,
    		lastMessage.time AS lastMessageTime,
    		lastMessage.authorid AS lastMessageAuthorid
		FROM chats
    		JOIN users ON (users.id <> $1 AND (users.id = chats.userid2 OR users.id = chats.userid1))
    		LEFT JOIN (
        		SELECT msg.text,
            		msg.time,
            		msg.messageOrder,
            		msg.chatid,
            		msg.authorid
        		FROM messages AS msg
        		WHERE msg.messageorder = (
                		SELECT MAX(messages2.messageOrder)
                		FROM messages AS messages2
                		WHERE msg.chatid = messages2.chatid
            		)
    		) lastMessage ON lastMessage.chatid = chats.id
		WHERE chats.userid1 = $1 OR chats.userid2 = $1
		ORDER BY lastMessageTime`,
		chatId, userId,
	)
	if err != nil {
		return model.Chat{}, err
	}

	err = repo.DB.Select(&chats[0].Photos, `SELECT photoId FROM photos WHERE userid = $1`, chats[0].PartnerId)
	if err != nil {
		fmt.Println(err)
		return model.Chat{}, err
	}

	return chats[0], nil
}
