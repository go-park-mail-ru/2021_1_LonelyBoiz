package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	model "server/internal/pkg/models"
	"time"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func (repo *MessageRepository) AddMessage(authorId int, chatId int, text string) (model.Message, error) {
	newMessage := model.Message{
		AuthorId: authorId,
		ChatId:   chatId,
		Text:     text,
		Time:     time.Now().Unix(),
		Reaction: -1,
	}

	err := repo.DB.QueryRow(
		`INSERT INTO messages (chatId, authorId, text, time, messageOrder)
			VALUES (
        	$1,
        	$2,
        	$3,
        	$4,
        	(
            	SELECT COUNT(*)
            	FROM messages
            	WHERE chatid = $1
        	) + 1
    	) RETURNING messageid, messageOrder`,
		newMessage.ChatId,
		newMessage.AuthorId,
		newMessage.Text,
		newMessage.Time,
	).Scan(&newMessage.MessageId, &newMessage.MessageOrder)
	if err != nil {
		return model.Message{}, err
	}

	return newMessage, nil
}

func (repo *MessageRepository) GetPartnerId(chatId int, userId int) (int, error) {
	var users []int
	err := repo.DB.Select(&users,
		`SELECT userid1, userid2 FROM chats WHERE id = $1`,
		chatId,
	)
	if err != nil {
		return -1, err
	}

	if users[0] != userId {
		return users[0], nil
	}

	return users[1], nil
}

func (repo *MessageRepository) GetMessages(chatId int, offset int, count int) ([]model.Message, error) {
	var messages []model.Message
	err := repo.DB.Select(&messages,
		`SELECT * FROM messages
			WHERE chatId = $1
			ORDER BY time
			LIMIT $2 OFFSET $3`,
		chatId,
		count,
		offset,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo *MessageRepository) ChangeMessage(messageId int, text string, reaction int) error {
	_, err := repo.DB.Exec(
		`UPDATE messages
			SET text = $1,
    		reaction = $2
		WHERE messageId = $3`,
		text,
		reaction,
		messageId,
	)

	return err
}

func (repo *MessageRepository) DeleteMessage(messageId int) error {
	_, err := repo.DB.Exec(
		`DELETE FROM messages WHERE messageid = $1`, messageId)

	return err
}
