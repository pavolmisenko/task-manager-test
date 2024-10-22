package models

import "errors"

type User struct {
	Id    string
	Name  string
	Email string
}

func (user *User) IsValid() error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Id == "" {
		return errors.New("failed to generate id")
	}
	return nil
}

func NewUser(newUser *User) (*User, error) {
	newUser.Id = GenerateId()
	err := newUser.IsValid()
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
