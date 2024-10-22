package models

import (
	"errors"
	"time"
)

type Task struct {
	Id          string
	Title       string
	Description string
	CreatedAt   time.Time
	Status      *Status
	Category    *Category
	User        *User
}

func (task *Task) isValid() error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

func NewTask(task *Task) (*Task, error) {
	task.Id = GenerateId()
	task.CreatedAt = time.Now()
	err := task.isValid()
	if err != nil {
		return nil, err
	}
	return task, nil
}
