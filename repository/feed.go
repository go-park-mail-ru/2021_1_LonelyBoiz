package repository

import (
	"database/sql"
	"server/api"

	_ "github.com/jackc/pgx/stdlib"
)

func (repo *RepoSqlx) CreateFeed(userId int) error {
	_, err := repo.DB.Exec(
		`INSERT into feed (
        SELECT user1.id,
            user2.id
        FROM users as user1
            inner join users user2 on (
                (
                    user1.datepreference = user2.sex
                    OR user1.datepreference = 'both'
                )
                AND (
                    user2.datepreference = user1.sex
                    OR user2.datepreference = 'both'
                )
                AND user1.id <> user2.id
				AND user2.isDeleted = false
				AND user2.idActive = true
            )
            AND user2.id NOT IN (
                SELECT userid2
                FROM feed
                WHERE userid1 = user1.id
            )
        WHERE user1.id = $1
    )`, userId)

	return err
}

func (repo *RepoSqlx) GetFeed(offset int) error {
	var users []api.User
	err := repo.DB.Select(&users,
		`SELECT * FROM feed
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

	return users, nil
}
