package repository

import (
	"main/database"
	"main/models"
	"time"
)

type NewPost struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	Url         string
	Enabled     bool
}

func newPost(description, url string) NewPost {
	return NewPost{
		Description: description,
		Url:         url,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Enabled:     true,
	}
}

type PostData struct {
	Description string
	Url         string
}

func CreatePost(data PostData) (int, error) {
	tx, err := database.Db.Begin()
	if err != nil {
		return 0, err
	}

	post := newPost(data.Description, data.Url)
	postId := 0
	err = tx.QueryRow(
		`INSERT INTO post (description, url, created_at, updated_at, enabled) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		post.Description, post.Url, post.CreatedAt, post.UpdatedAt, post.Enabled,
	).Scan(&postId)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return postId, nil
}

type GetPostsParams struct {
	Take int
	Skip int
}

func GetPosts(params GetPostsParams) ([]models.Post, error) {
	rows, err := database.Db.Query(
		`SELECT id, description, url, created_at, updated_at, enabled FROM post WHERE enabled = true LIMIT $1 OFFSET $2`,
		params.Take, params.Skip,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Description, &post.Url, &post.CreatedAt, &post.UpdatedAt, &post.Enabled)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
