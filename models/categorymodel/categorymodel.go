package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`
		SELECT id, name, updated_at
		FROM categories
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.UpdateAt,
		)
		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
