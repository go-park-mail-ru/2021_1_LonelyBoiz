package repository

import (
	"github.com/jmoiron/sqlx"
	model "server/internal/pkg/models"
)

type ChatRepository struct {
	DB *sqlx.DB
}

func (repo *ChatRepository) GetChats(userId int, limit int, offset int) ([]model.Chat, error) {
	var chats []model.Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id AS chatId,
    		users.id AS partnerId,
    		users.name AS partnerName,
    		lastMessage.text AS lastMessage,
    		lastMessage.time AS lastMessageTime,
    		lastMessage.authorid AS lastMessageAuthorid,
    		users.photos AS avatar
		FROM chats
    		JOIN users ON ((users.id = chats.userid2 OR users.id = chats.userid1) AND users.id <> $1)
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
		ORDER BY lastMessageTime
		LIMIT $2 OFFSET $3`,
		userId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (repo *ChatRepository) GetChat(chatId int, limit int, offset int) ([]model.Message, error) {
	var messages []model.Message
	err := repo.DB.Select(&messages,
		`SELECT * FROM messages
			WHERE messages.chatid = $1
			ORDER BY messages.messageorder
			LIMIT $2 OFFSET $3;`,
		chatId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo *ChatRepository) CreateChat(userId1 int, userId2 int) (int, error) {
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

func (repo *ChatRepository) GetPartnerId(chatId int, userId int) (int, error) {
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
