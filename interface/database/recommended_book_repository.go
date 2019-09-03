package database

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type RecommendedBookRepository struct {
	SqlHandler
}

func (repo *RecommendedBookRepository) All(limit int) (recommendedBooks model.RecommendedBooks, err error) {
	query := "SELECT id, link, image_url, button_url FROM recommended_books"
	query = query + " ORDER BY id DESC LIMIT ?"
	rows, err := repo.SqlHandler.Query(query, limit)
	if err != nil {
		return recommendedBooks, err
	}
	defer rows.Close()

	for rows.Next() {
		r := model.RecommendedBook{}
		if err := rows.Scan(&r.Id, &r.Link, &r.ImageUrl, &r.ButtonUrl); err != nil {
			return recommendedBooks, err
		}

		recommendedBooks = append(recommendedBooks, r)
	}
	return
}
