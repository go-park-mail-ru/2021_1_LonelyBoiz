package repository

import (
	"database/sql"
	"errors"
	"fmt"
	model "server/internal/pkg/models"

	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type UserRepositoryInterface interface {
	AddUser(newUser model.User) (int, error)
	GetUser(id int) (model.User, error)
	DeleteUser(id int) error
	ChangeUser(newUser model.User) error
	CheckMail(email string) (bool, error)
	GetPassWithEmail(email string) ([]byte, error)
	GetPassWithId(id int) ([]byte, error)
	SignIn(email string) (model.User, error)
	ChangePassword(userId int, hash []byte) error

	//фотки
	GetPhotos(userId int) ([]string, error)

	//чат
	CreateChat(userId1 int, userId2 int) (int, error)
	GetNewChatById(chatId int, userId int) (model.Chat, error)
	GetChatById(chatId int, userId int) (model.Chat, error)

	//лента
	ClearFeed(userId int) error
	CreateFeed(userId int) error
	GetFeed(userId int, limit int) ([]int, error)

	//уеньшить количесвто скроллов
	ReduceScrolls(userId int) (int, error)

	//реакция
	Rating(userIdFrom int, userIdTo int, reaction string) (int64, error)
	CheckReciprocity(userId1 int, userId2 int) (bool, error)
	//добавить скролов в ленту
	UpdatePayment(userId int, amount int) error

	// секретный альбом
	UnblockSecreteAlbum(ownerId int, getterId int) error
	CheckPermission(ownerId int, getterId int) (bool, error)
	GetSecretePhotos(ownerId int) ([]string, error)
	AddToSecreteAlbum(ownerId int, photos []string) error
}

type UserRepository struct {
	DB *sqlx.DB
}

func (repo *UserRepository) ReduceScrolls(userId int) (int, error) {
	var amount int
	err := repo.DB.QueryRowx(
		`UPDATE users SET scrolls = (scrolls - 1)
  			WHERE id = $1
  		RETURNING scrolls;`,
		userId,
	).Scan(&amount)
	if err != nil {
		return -1, err
	}

	return amount, nil
}

func (repo *UserRepository) UpdatePayment(userId int, amount int) error {
	_, err := repo.DB.Exec(
		`UPDATE users
			SET scrolls = case
        		when scrolls < 1 then $1
        		when scrolls > 0 then scrolls + $1
    		end
		WHERE id = $2;`,
		amount, userId,
	)

	return err
}

func (repo *UserRepository) AddToSecreteAlbum(ownerId int, photos []string) error {
	err := repo.DB.QueryRowx(
		`INSERT INTO secretPhotos (userId, photos) Values ($1, $2)`,
		ownerId, pq.Array(photos))
	return err.Err()
}

func (repo *UserRepository) GetSecretePhotos(ownerId int) ([]string, error) {
	var photos []pq.StringArray
	err := repo.DB.Select(&photos,
		`SELECT photos FROM secretPhotos WHERE userId = $1`,
		ownerId)
	if err != nil {
		return nil, err
	}
	if len(photos) == 0 {
		return make([]string, 0), nil
	}

	return photos[0], nil
}

func (repo *UserRepository) CheckPermission(ownerId int, getterId int) (bool, error) {
	var ids []int
	err := repo.DB.Select(&ids,
		`SELECT ownerId FROM secretPermission WHERE ownerId = $1 AND getterId = $2`,
		ownerId, getterId)
	if err != nil {
		return false, err
	}
	if len(ids) == 0 {
		return false, nil
	}

	return true, nil
}

func (repo *UserRepository) UnblockSecreteAlbum(ownerId int, getterId int) error {
	err := repo.DB.QueryRowx(
		`INSERT INTO secretPermition (ownerId, getterId) Values ($1, $2)`,
		ownerId, getterId)
	return err.Err()
}

func (repo *UserRepository) AddUser(newUser model.User) (int, error) {
	var id int
	newUser.Photos = make([]string, 0)

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
			instagram,
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
			$11,
			$12
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
		pq.Array(newUser.Photos),
	).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repo *UserRepository) GetUser(id int) (model.User, error) {
	var user []model.User
	fmt.Println("here")
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
    		isdeleted,
			photos
		FROM users WHERE id = $1`,
		id)
	if err != nil {
		return model.User{}, err
	}
	if len(user) == 0 {
		return model.User{}, sql.ErrNoRows
	}

	if len(user[0].Photos) == 0 {
		user[0].Photos = make([]string, 0)
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
			datePreference = $7, isActive = $8, instagram = $9, photos = $10
		WHERE id = $11`,
		newUser.Email, newUser.Name, newUser.Birthday,
		newUser.Description, newUser.City, newUser.Sex,
		newUser.DatePreference, newUser.IsActive, newUser.Instagram, newUser.Photos,
		newUser.Id,
	)

	return err
}

func (repo *UserRepository) CheckMail(email string) (bool, error) {
	var emails []string
	err := repo.DB.Select(&emails, `SELECT email FROM users WHERE email = $1`, email)
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
    		isdeleted,
			photos
		FROM users WHERE email = $1`, email)
	if err != nil {
		return model.User{}, err
	}
	if len(user) == 0 {
		return model.User{}, errors.New("пользователь не найден")
	}

	return user[0], nil
}

func (repo *UserRepository) GetPhotos(userId int) ([]string, error) {
	var photos pq.StringArray
	err := repo.DB.Select(&photos, `SELECT photos FROM users WHERE Id = $1`, userId)
	if err != nil {
		return nil, err
	}
	if len(photos) == 0 {
		return make([]string, 0), nil
	}

	return photos, nil
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
    	) RETURNING id`,
		userId1, userId2,
	).Scan(&chatId)
	if err != nil {
		return -1, err
	}

	return chatId, err
}

func (repo *UserRepository) GetNewChatById(chatId int, userId int) (model.Chat, error) {
	var chats []model.Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id AS chatId,
    		users.id AS partnerId,
    		users.name AS partnerName
		FROM chats
    		JOIN users ON (users.id <> $1 AND (users.id = chats.userid2 OR users.id = chats.userid1))
		WHERE (chats.userid1 = $1 OR chats.userid2 = $1) AND chats.id = $2`,
		userId, chatId,
	)
	if err != nil {
		return model.Chat{}, err
	}
	if len(chats) == 0 {
		return model.Chat{}, nil
	}

	err = repo.DB.Select(&chats[0].Photos, `SELECT photos FROM users WHERE id = $1`, chats[0].PartnerId)
	if err != nil {
		return model.Chat{}, err
	}

	return chats[0], nil

}

func (repo *UserRepository) GetChatById(chatId int, userId int) (model.Chat, error) {
	var chats []model.Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id AS chatid,
    		users.id AS partnerid,
    		users.name AS partnername,
			users.photos AS photos,
    		lastMessage.text AS lastmessage,
    		lastMessage.time AS lastmessagetime,
    		lastMessage.authorid AS lastmessageauthorid
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
	if len(chats) == 0 {
		return model.Chat{}, nil
	}

	return chats[0], nil
}
