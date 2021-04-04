package repository

import (
	_ "github.com/jackc/pgx/stdlib"
)

type Chat struct {
	ChatId              int    `json:"chatId"`
	PartnerId           int    `json:"partnerId"`
	PartnerName         string `json:"partnerName"`
	LastMessage         string `json:"lastMessage"`
	LastMessageTime     int64  `json:"lastMessageTime"`
	LastMessageAuthorId int    `json:"lastMessageAuthor"`
	Avatar              string `json:"pathToAvatar"`
}

func (repo *RepoSqlx) GetChats(userId int, limit int, offset int) ([]Chat, error) {
	var chats []Chat
	err := repo.DB.Select(&chats,
		`SELECT chats.id as chatId,
    		chats.userid2 as partnerId,
    		users.name as partnerName,
    		lastMessage.text as lastMessage,
    		lastMessage.time as lastMessageTime,
    		lastMessage.authorid as lastMessageAuthorid,
    		users.photos as avatar
		FROM chats
    		join users on (users.id = chats.userid2)
    		join (
        		SELECT msg.text,
            		msg.time,
            		msg.messageOrder,
            		msg.chatid,
            		msg.authorid
        		FROM messages as msg
        		where msg.messageorder = (
                		select max(messages2.messageOrder)
                		from messages as messages2
                		where msg.chatid = messages2.chatid
            		)
    		) lastMessage on lastMessage.chatid = chats.id
		where chats.userid1 = $1
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
