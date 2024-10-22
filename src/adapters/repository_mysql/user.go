package repository_mysql

import (
	"database/sql"
	"taskmanager/src/models"
)

type User struct {
	Pool *sql.DB
}

func (u User) GetAll() ([]*models.User, error) {
	stmt, err := u.Pool.Prepare(`
	SELECT u.id AS 		userId,
		   u.name AS 	userName,
		   u.email AS 	userEmail
	FROM User u
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {
		var userId, userName, userEmail string
		err := rows.Scan(
			&userId,
			&userName,
			&userEmail,
		)
		if err != nil {
			return nil, err
		}
		user := models.User{
			Id:    userId,
			Name:  userName,
			Email: userEmail,
		}
		users = append(users, &user)
	}
	defer rows.Close()

	return users, nil

}

func (u User) GetById(id string) (*models.User, error) {
	// TBD
	return nil, nil
}

func (u User) Create(user *models.User) (*models.User, error) {
	stmt, err := u.Pool.Prepare(`
	INSERT INTO User (id, name, email)
	VALUES (?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u User) Delete(id string) error {
	// Start a transaction
	tx, err := u.Pool.Begin()
	if err != nil {
		return err
	}

	// Delete tasks associated with the user
	_, err = tx.Exec(`
	DELETE FROM Tasks
	WHERE user_id = ?
	`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete the user
	_, err = tx.Exec(`
	DELETE FROM User
	WHERE id = ?
	`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

