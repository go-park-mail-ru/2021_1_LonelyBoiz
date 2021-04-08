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

func (repo *RepoSqlx) GetChat(chatId int, limit int, offset int) ([]Message, error) {
	var messages []Message
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

func (repo *RepoSqlx) CreateChat(userId1 int, userId2 int) (int, error) {
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
