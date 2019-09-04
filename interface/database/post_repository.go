package database

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"time"
)

type PostRepository struct {
	SqlHandler
}

func (repo *PostRepository) Index(page int) (posts model.Posts, err error) {
	query := "SELECT id, post_category_id, title, image_file_name, published_at FROM posts WHERE active = 1 AND published_at <= ? ORDER BY published_at DESC"
	query = query + " LIMIT 10 OFFSET ?"
	offset := 10 * (page - 1)
	nowTime := time.Now()
	rows, err := repo.SqlHandler.Query(query, nowTime, offset)
	defer rows.Close()

	if err != nil {
		return posts, err
	}

	for rows.Next() {
		p := model.Post{}
		err := rows.Scan(&p.Id, &p.PostCategoryId, &p.Title, &p.ImageUrl, &p.PublishedAt)
		if err != nil {
			return posts, err
		}

		posts = append(posts, p)
	}
	return
}

func (repo *PostRepository) FindById(id int) (post model.Post, err error) {
	nowTime := time.Now()
	query := "SELECT id, post_category_id, title, content, image_file_name, published_at FROM posts WHERE id = ? AND active = 1 AND published_at <= ? LIMIT 1"
	rows, err := repo.SqlHandler.Query(query, id, nowTime)
	if err != nil {
		return post, err
	}

	for rows.Next() {
		err := rows.Scan(&post.Id, &post.PostCategoryId, &post.Title, &post.Content, &post.ImageUrl, &post.PublishedAt)
		if err != nil {
			return post, err
		}
		break
	}
	return
}

func (repo *PostRepository) GetPostsCount() (count int, err error) {
	nowTime := time.Now()
	query := "SELECT COUNT(*) FROM posts WHERE active = 1 AND published_at <= ? LIMIT 1"
	rows, err := repo.SqlHandler.Query(query, nowTime)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
		break
	}
	return
}
