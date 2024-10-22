package repository_mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"taskmanager/src/models"
)

type Task struct {
	Pool *sql.DB
}

func ReconstructTask(
	taskId string,
	taskTitle string,
	taskDescription string,
	taskCreatedAt string,
	taskStatusId string,
	taskCategoryId string,
	userId string,
	userName string,
	userEmail string,
	categoryId string,
	categoryName string,
	statusId string,
	statusName string,
) (*models.Task, error) {
	parsedTaskCreatedAt, err := StringToTime(taskCreatedAt)
	if err != nil {
		return nil, errors.New("Failed to parse data")
	}
	status := models.Status{
		Id:   statusId,
		Name: statusName,
	}
	category := models.Category{
		Id:   categoryId,
		Name: categoryName,
	}
	user := models.User{
		Id:    userId,
		Name:  userName,
		Email: userEmail,
	}
	return &models.Task{
		Id:          taskId,
		Title:       taskTitle,
		Description: taskDescription,
		CreatedAt:   parsedTaskCreatedAt,
		Status:      &status,
		Category:    &category,
		User:        &user,
	}, nil
}

func (t Task) GetAll() ([]*models.Task, error) {
	stmt, err := t.Pool.Prepare(`
	SELECT t.id AS 			taskId,
		   t.title AS 		taskTitle,
		   t.description AS taskDescription,
		   t.created_at AS 	taskCreatedAt,
		   t.status_id AS 	taskStatusId,
		   t.category_id AS taskCategoryId,
		   u.id AS 			userId,
		   u.name AS 		userName,
		   u.email AS 		userEmail,
		   c.id AS 			categoryId,
		   c.name AS 		categoryName,
		   s.id AS 			statusId,
		   s.name AS 		statusName
	FROM Tasks t
	LEFT JOIN Status s ON t.status_id = s.id
	LEFT JOIN Category c ON t.category_id = c.id
	LEFT JOIN User u ON t.user_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		var taskId, taskTitle, taskDescription, taskCreatedAt, taskStatusId, taskCategoryId, userId, userName, userEmail, categoryId, categoryName, statusId, statusName string
		err := rows.Scan(
			&taskId,
			&taskTitle,
			&taskDescription,
			&taskCreatedAt,
			&taskStatusId,
			&taskCategoryId,
			&userId,
			&userName,
			&userEmail,
			&categoryId,
			&categoryName,
			&statusId,
			&statusName,
		)
		if err != nil {
			return nil, err
		}
		task, err := ReconstructTask(
			taskId,
			taskTitle,
			taskDescription,
			taskCreatedAt,
			taskStatusId,
			taskCategoryId,
			userId,
			userName,
			userEmail,
			categoryId,
			categoryName,
			statusId,
			statusName,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t Task) GetById(id string) (*models.Task, error) {
	stmt, err := t.Pool.Prepare(`
	SELECT t.id AS 			taskId,
		   t.title AS 		taskTitle,
		   t.description AS taskDescription,
		   t.created_at AS 	taskCreatedAt,
		   t.status_id AS 	taskStatusId,
		   t.category_id AS taskCategoryId,
		   u.id AS 			userId,
		   u.name AS 		userName,
		   u.email AS 		userEmail,
		   c.id AS 			categoryId,
		   c.name AS 		categoryName,
		   s.id AS 			statusId,
		   s.name AS 		statusName
	FROM Tasks t
	JOIN Status s ON t.status_id = s.id
	JOIN Category c ON t.category_id = c.id
	JOIN User u ON t.user_id = u.id
	WHERE t.id = ?
	`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var taskId, taskTitle, taskDescription, taskCreatedAt, taskStatusId, taskCategoryId, userId, userName, userEmail, categoryId, categoryName, statusId, statusName string
	err = stmt.QueryRow(id).Scan(
		&taskId,
		&taskTitle,
		&taskDescription,
		&taskCreatedAt,
		&taskStatusId,
		&taskCategoryId,
		&userId,
		&userName,
		&userEmail,
		&categoryId,
		&categoryName,
		&statusId,
		&statusName,
	)
	if err != nil {
		return nil, err
	}
	task, err := ReconstructTask(
		taskId,
		taskTitle,
		taskDescription,
		taskCreatedAt,
		taskStatusId,
		taskCategoryId,
		userId,
		userName,
		userEmail,
		categoryId,
		categoryName,
		statusId,
		statusName,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t Task) Create(task *models.Task) (*models.Task, error) {
	smtm, err := t.Pool.Prepare(`
	INSERT INTO Tasks (id, title, description, created_at, status_id, category_id, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer smtm.Close()
	fmt.Println("repository", task)
	_, err = smtm.Exec(
		task.Id,
		task.Title,
		task.Description,
		task.CreatedAt,
		task.Status.Id,
		task.Category.Id,
		task.User.Id,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t Task) Delete(id string) error {
    stmt, err := t.Pool.Prepare(`
    DELETE FROM Tasks WHERE id = ?
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the statement and check for errors
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    // Return nil if everything went ok
    return nil
}
