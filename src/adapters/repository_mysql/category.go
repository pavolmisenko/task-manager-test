package repository_mysql

import (
	"database/sql"
	"taskmanager/src/models"
)

type Category struct {
	Pool *sql.DB
}

func (c Category) GetAll() ([]*models.Category, error) {
	stmt, err := c.Pool.Prepare(`
	SELECT c.id AS 		categoryId,
		   c.name AS 	categoryName
	FROM Category c
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	categories := []*models.Category{}

	for rows.Next() {
		var categoryId, categoryName string
		err := rows.Scan(
			&categoryId,
			&categoryName,
		)
		if err != nil {
			return nil, err
		}
		category := models.Category{
			Id:   categoryId,
			Name: categoryName,
		}
		categories = append(categories, &category)
	}
	defer rows.Close()

	return categories, nil
}

func (c Category) GetById(id string) (*models.Category, error) {
	// TBD
	return nil, nil
}
