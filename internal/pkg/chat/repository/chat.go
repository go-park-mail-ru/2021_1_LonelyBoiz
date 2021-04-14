package repository

import (
	"fmt"
	model "server/internal/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	DB *sqlx.DB
}

func reverseChats(chats []model.Chat) []model.Chat {
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
    		COALESCE(lastMessage.text, '') AS lastMessage,
    		COALESCE(lastMessage.time, 0) lastMessage.time AS lastMessageTime,
    		COALESCE(lastMessage.authorid, -1) AS lastMessageAuthorid
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
	if len(chats) == 0 {
		chats = make([]model.Chat, 0)
		return chats, nil
	}

	for i, _ := range chats {
		err = repo.DB.Select(&chats[i].Photos, `SELECT photoId FROM photos WHERE userid = $1`, chats[i].PartnerId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if len(chats[i].Photos) == 0 {
			chats[i].Photos = make([]int, 0)
		}
	}

	return reverseChats(chats), nil
}
