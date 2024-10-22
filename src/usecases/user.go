package usecases

import (
	"taskmanager/src/models"
	"taskmanager/src/repository"
)

type NewUser struct {
	Name  string
	Email string
}

type IUser interface {
	GetAll() ([]*models.User, error)
	GetById(id string) (*models.User, error)
	Create(user *NewUser) (*models.User, error)
	Delete(id string) error
}

type User struct {
	Repo repository.IUser
}

func (u User) GetAll() ([]*models.User, error) {
	return u.Repo.GetAll()
}

func (u User) GetById(id string) (*models.User, error) {
	return u.Repo.GetById(id)
}

func (u User) Create(user *NewUser) (*models.User, error) {
	// Create a new user
	newUser, err := models.NewUser(
		&models.User{
			Name:  user.Name,
			Email: user.Email,
		},
	)
	if err != nil {
		return nil, err
	}
	return u.Repo.Create(newUser)
}

func (u User) Delete(id string) error {
	return u.Repo.Delete(id)
}

