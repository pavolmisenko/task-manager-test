package repository_mysql

import (
	"database/sql"
	"taskmanager/src/models"
)

type Status struct {
	Pool *sql.DB
}

func (s Status) GetAll() ([]*models.Status, error) {
	stmt, err := s.Pool.Prepare(`
	SELECT s.id AS 		statusId,
		   s.name AS 	statusName
	FROM Status s
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	statuses := []*models.Status{}

	for rows.Next() {
		var statusId, statusName string
		err := rows.Scan(
			&statusId,
			&statusName,
		)
		if err != nil {
			return nil, err
		}
		status := models.Status{
			Id:   statusId,
			Name: statusName,
		}
		statuses = append(statuses, &status)
	}
	defer rows.Close()

	return statuses, nil
}
func (s Status) GetById(id string) (*models.Status, error) {
	// TBD
	return nil, nil
}
