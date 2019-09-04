package database

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type PostCategoryRepository struct {
	SqlHandler
}

func (repo *PostCategoryRepository) FindById(id int) (postCategory model.PostCategory, err error) {
	query := "SELECT id, name, color FROM post_categories WHERE id = ?"
	rows, err := repo.SqlHandler.Query(query, id)
	if err != nil {
		return postCategory, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&postCategory.Id, &postCategory.Name, &postCategory.Color)
		if err != nil {
			return postCategory, err
		}
		break
	}
	return
}
