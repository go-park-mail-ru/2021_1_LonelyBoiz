package repository

import (
	model "server/internal/pkg/models"

	"github.com/jmoiron/sqlx"
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
