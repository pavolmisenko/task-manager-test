package usecases

import (
	"taskmanager/src/models"
	"taskmanager/src/repository"
)

type NewTask struct {
	Title			string
	Description		string
	UserId			string
	CategoryId		string
	StatusId		string
}

type ITask interface {
	GetAll() ([]*models.Task, error)
	GetById(id string) (*models.Task, error)
	Create(task *NewTask) (*models.Task, error)
	Delete(id string) error
}

type Task struct {
	Repo repository.ITask
}

func (t Task) GetAll() ([]*models.Task, error) {
	return t.Repo.GetAll()
}

func (t Task) GetById(id string) (*models.Task, error) {
	return t.Repo.GetById(id)
}

func (t Task) Create(task *NewTask) (*models.Task, error) {
	newTask, err := models.NewTask(&models.Task{
		Title: task.Title,
		Description: task.Description,
		User: &models.User{Id: task.UserId},
		Category: &models.Category{Id: task.CategoryId},
		Status: &models.Status{Id: task.StatusId},
	})
	if err != nil {
		return nil, err
	}
	return t.Repo.Create(newTask)
}

func (t Task) Delete(id string) error {
	return t.Repo.Delete(id)
}

