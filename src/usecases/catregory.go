package usecases

import (
	"taskmanager/src/models"
	"taskmanager/src/repository"
)

type ICategory interface {
	GetAll() ([]*models.Category, error)
	GetById(id string) (*models.Category, error)
}

type Category struct {
	Repo repository.ICategory
}

func (c Category) GetAll() ([]*models.Category, error) {
	return c.Repo.GetAll()
}

func (c Category) GetById(id string) (*models.Category, error) {
	return c.Repo.GetById(id)
}
