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
        WHERE user1.id = $1
        LIMIT 100
    )`, userId)

	return err
}

func (repo *RepoSqlx) GetFeed(userId int) ([]api.User, error) {
	var feed []api.User
	err := repo.DB.Select(&feed,
		`SELECT users.id,
    		users.email,
    		users.name,
    		users.city,
    		users.sex,
			users.birthday,
    		users.datePreference,
    		users.city,
    		users.description,
			users.photos,
			users.isActive,
			users.isDeleted
    	FROM feed
    		join users on userid2 = users.id
		WHERE userid1 = $1 AND rating = 'empty' LIMIT 20`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (repo *RepoSqlx) Rating(userIdFrom int, userIdTo int, reaction string) (int64, error) {
	res, err := repo.DB.Exec(
		`UPDATE feed
			SET rating = $1
		WHERE userid1 = $2 AND userid2 = $3`,
		reaction, userIdFrom, userIdTo,
	)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (repo *RepoSqlx) CheckReciprocity(userId1 int, userId2 int) (bool, error) {
	var rating []string
	err := repo.DB.Select(&rating,
		`SELECT rating
			FROM feed
			WHERE userid1 = $1 AND userid2 = $2`,
		userId1, userId2,
	)
	if err != nil {
		return false, err
	}

	if len(rating) != 0 && rating[0] == "like" {
		return true, nil
	}

	return false, nil
}