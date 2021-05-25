package repository

import (
	model "server/internal/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ChatRepositoryInterface interface {
	GetChats(userId int, limit int, offset int) ([]model.Chat, error)
}

type ChatRepository struct {
	DB *sqlx.DB
}

func (repo *ChatRepository) reverseChats(chats []model.Chat) []model.Chat {
	newChats := make([]model.Chat, 0, len(chats))
	for i := len(chats) - 1; i >= 0; i-- {
		newChats = append(newChats, chats[i])
	}
	return newChats
}

func (repo *ChatRepository) GetChats(userId int, limit int, offset int) ([]model.Chat, error) {
	var chats []model.Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id AS chatId,
    		users.id AS partnerId,
    		users.name AS partnerName,
			users.photos AS photos,
    		COALESCE(lastMessage.text, '') AS lastMessage,
    		COALESCE(lastMessage.time, 0) AS lastMessageTime,
    		COALESCE(lastMessage.authorid, -1) AS lastMessageAuthorid,
			CASE when secretPermission.getterId IS NULL then FALSE
        	    ELSE TRUE
                END as isOpened
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
			LEFT JOIN secretPermission on (ownerId = $1 AND users.id = getterId)
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
	if len(chats) == 0 {
		chats = make([]model.Chat, 0)
		return chats, nil
	}

	return repo.reverseChats(chats), nil
}
