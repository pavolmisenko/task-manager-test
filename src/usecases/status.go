package usecases

import (
	"taskmanager/src/models"
	"taskmanager/src/repository"
)

type IStatus interface {
	GetAll() ([]*models.Status, error)
	GetById(id string) (*models.Status, error)
}

type Status struct {
	Repo repository.IStatus
}

func (s Status) GetAll() ([]*models.Status, error) {
	return s.Repo.GetAll()
}

func (s Status) GetById(id string) (*models.Status, error) {
	return s.Repo.GetById(id)
}
