package repository

import "taskmanager/src/models"

type ICategory interface {
	GetAll() ([]*models.Category, error)
	GetById(id string) (*models.Category, error)
}

type IStatus interface {
	GetAll() ([]*models.Status, error)
	GetById(id string) (*models.Status, error)
}

type IUser interface {
	GetAll() ([]*models.User, error)
	GetById(id string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Delete(id string) error
}

type ITask interface {
	GetAll() ([]*models.Task, error)
	GetById(id string) (*models.Task, error)
	Create(task *models.Task) (*models.Task, error)
	Delete(id string) error
}
