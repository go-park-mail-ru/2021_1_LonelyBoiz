package repository

import (
	model "server/internal/pkg/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func (repo *MessageRepository) CheckMessageForReacting(userId int, messageId int) (int, error) {
	var ids []int
	err := repo.DB.Select(&ids,
		`SELECT authorId FROM messages
    		JOIN chats ON 
				messages.chatid = chats.id
    		AND(
        		chats.userid1 = $1 OR chats.userid2 = $1
    		)
			WHERE messages.messageid = $2;`,
		userId, messageId)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}

	return ids[0], nil
}

func (repo *MessageRepository) CheckChat(userId int, chatId int) (bool, error) {
	var ids []int
	err := repo.DB.Select(&ids,
		`SELECT id FROM chats 
			WHERE (userid1 = $1 OR userid2 = $1) 
				AND (id = $2)`,
		userId, chatId)
	if err != nil {
		return false, err
	}
	if len(ids) == 0 {
		return false, nil
	}

	return true, nil
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
		`SELECT CASE
				WHEN (userid1 = $1) THEN userid2
				WHEN (userid2 = $1) THEN userid1
			END
			FROM chats
			WHERE id = $2`,
		userId, chatId,
	)
	if err != nil {
		return -1, err
	}
	if len(users) == 0 {
		return -1, nil
	}

	return users[0], nil
}

func (repo *MessageRepository) ChangeMessageText(messageId int, text string) error {
	_, err := repo.DB.Exec(
		`UPDATE messages
			SET text = $1
		WHERE messageId = $2`,
		text,
		messageId,
	)

	return err
}

func (repo *MessageRepository) ChangeMessageReaction(messageId int, reaction int) error {
	_, err := repo.DB.Exec(
		`UPDATE messages
			SET reaction = $1
		WHERE messageId = $2`,
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

func reverseMessages(messages []model.Message) []model.Message {
	newMessages := make([]model.Message, 0, len(messages))
	for i := len(messages) - 1; i >= 0; i-- {
		newMessages = append(newMessages, messages[i])
	}
	return newMessages
}

func (repo *MessageRepository) GetMessages(chatId int, limit int, offset int) ([]model.Message, error) {
	var messages []model.Message
	err := repo.DB.Select(&messages,
		`SELECT * FROM messages
			WHERE messages.chatid = $1
			ORDER BY messages.messageorder DESC
			LIMIT $2 OFFSET $3`,
		chatId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return reverseMessages(messages), nil
}
