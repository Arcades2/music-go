package repository

import (
	"main/database"
	"time"
)

type NewLike struct {
	CreatedAt time.Time
	PostId    int
}

func newLike(postId int) NewLike {
	return NewLike{
		PostId:    postId,
		CreatedAt: time.Now(),
	}
}

type LikeData struct {
	PostId int
}

func CreateLike(data LikeData) (int, error) {
	tx, err := database.Db.Begin()
	if err != nil {
		return 0, err
	}

	like := newLike(data.PostId)
	likeId := 0
	err = tx.QueryRow(
		`INSERT INTO "like" (post_id, created_at) VALUES ($1, $2) RETURNING id`,
		like.PostId, like.CreatedAt,
	).Scan(&likeId)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return likeId, nil
}
