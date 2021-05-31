package repository

func (repo *UserRepository) CleanFeed(userId int) error {
	_, err := repo.DB.Exec("DELETE FROM feed WHERE userid1 = $1 AND (rating = 'skip' OR rating = 'empty')", userId)

	return err
}

func (repo *UserRepository) ClearFeed(userId int) error {
	_, err := repo.DB.Exec(
		`UPDATE feed SET rating = 'empty' WHERE userid1 = $1 AND rating = 'skip'`,
		userId,
	)

	return err
}

func (repo *UserRepository) CreateFeed(userId int) error {
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
				AND (
					user1.partnerHeightTop = -1
					OR (user1.partnerHeightTop >= user2.height AND user1.partnerHeightBot <= user2.height)
				)
				AND (
					user2.partnerHeightTop = -1
					OR (user2.partnerHeightTop >= user1.height AND user2.partnerHeightBot <= user1.height)
				)
				AND (
					user1.partnerWeightTop = -1
					OR (user1.partnerWeightTop >= user2.weight AND user1.partnerWeightBot <= user2.weight)
				)
				AND (
					user2.partnerWeightTop = -1
					OR (user2.partnerWeightTop >= user1.weight AND user2.partnerWeightBot <= user1.weight)
				)
				AND (
					user1.partnerAgeTop = -1
					OR (user1.partnerAgeTop >= (user2.birthday/60/60/24/365-2003+1970) AND user1.partnerAgeBot <= (user2.birthday/60/60/24/365-2003+1970))
				)
				AND (
					user2.partnerAgeTop = -1
					OR (user2.partnerAgeTop >= (user1.birthday/60/60/24/365-2003+1970) AND user2.partnerAgeBot <= (user1.birthday/60/60/24/365-2003+1970))
				)
				AND (
					(array_length(user1.interests, 1) = 0) IS NULL
					OR (user1.interests && user2.interests)
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

func (repo *UserRepository) GetFeed(userId int, limit int) ([]int, error) {
	var feed []int
	err := repo.DB.Select(&feed,
		`SELECT userid2
    		FROM feed
			WHERE userid1 = $1 AND rating = 'empty' LIMIT $2`,
		userId, limit,
	)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (repo *UserRepository) Rating(userIdFrom int, userIdTo int, reaction string) (int64, error) {
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

func (repo *UserRepository) CheckReciprocity(userId1 int, userId2 int) (bool, error) {
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
