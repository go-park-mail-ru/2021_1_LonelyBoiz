package repository

import (
	"server/api"

	_ "github.com/jackc/pgx/stdlib"
)

func (repo *RepoSqlx) CreateFeed(userId int) error {
	_, err := repo.DB.Exec(
		`INSERT into feed (
        SELECT user1.id,
            user2.id
        FROM users as user1
            JOIN users user2 on (
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
                AND user2.isActive = true
                AND user2.id NOT IN (
                    SELECT userid2
                    FROM feed
                    WHERE userid1 = user1.id
                )
            )
        WHERE user1.id = 1
        LIMIT 100
    )`, userId)

	return err
}

func (repo *RepoSqlx) GetFeed(userId int) ([]api.User, error) {
	var feed []api.User
	err := repo.DB.Select(&feed,
		`SELECT *
			FROM feed
    			join users on userid2 = users.id
			WHERE userid1 = 1 AND rating = 'empty' LIMIT 20`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
