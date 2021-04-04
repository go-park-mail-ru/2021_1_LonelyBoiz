package repository

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

type Message struct {
	MessageId    int    `json:"messageId"`
	AuthorId     int    `json:"authorId"`
	ChatId       int    `json:"chatId"`
	Text         string `json:"text"`
	Reaction     int    `json:"reaction"`
	Time         int64  `json:"time"`
	MessageOrder int    `json:"messageOrder"`
}

func (repo *RepoSqlx) AddMessage(authorId int, chatId int, text string) (Message, error) {
	newMessage := Message{
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
            	WHERE chatid = $2
        	) + 1
    	)`,
		newMessage.ChatId,
		newMessage.AuthorId,
		newMessage.Text,
		newMessage.Time,
		newMessage.Reaction,
	).Scan(&newMessage.MessageId)
	if err != nil {
		return Message{}, err
	}

	return newMessage, nil
}

func (repo *RepoSqlx) GetMessages(chatId int, offset int, count int) ([]Message, error) {
	var messages []Message
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

func (repo *RepoSqlx) ChangeMessage(messageId int, text string, reaction int) error {
	_, err := repo.DB.Exec(
		`UPDATE messages
			SET text = $1,
    		reaction = $2
		WHERE messageId = #3`,
		text,
		reaction,
		messageId,
	)

	return err
}

func (repo *RepoSqlx) DeleteMessage(messageId int) error {
	_, err := repo.DB.Exec(
		`DELETE FROM messages WHERE messageid = $1`, messageId)

	return err
}
