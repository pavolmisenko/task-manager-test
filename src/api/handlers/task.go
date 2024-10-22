package handlers

import (
	"fmt"
	"net/http"
	"taskmanager/src/models"
	"taskmanager/src/usecases"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	TaskUsecase     usecases.ITask
	CategoryUsecase usecases.ICategory
	StatusUsecase   usecases.IStatus
	UserUsecase     usecases.IUser
}

// Handler for form data submission.
func (h TaskHandler) BaseHandler(context echo.Context) error {
	categories, err := h.CategoryUsecase.GetAll()
	if err != nil {
		return err
	}
	statuses, err := h.StatusUsecase.GetAll()
	if err != nil {
		return err
	}
	users, err := h.UserUsecase.GetAll()
	if err != nil {
		return err
	}
	darkMode := false
    if cookie, err := context.Cookie("dark-mode"); err == nil {
        darkMode = cookie.Value == "true"
    }
	data := struct {
        Categories []*models.Category
        Users      []*models.User
        Statuses   []*models.Status
        DarkMode   bool
    }{
        Categories: categories,
        Users:      users,
        Statuses:   statuses,
        DarkMode:   darkMode,
    }
	return context.Render(http.StatusOK, "tasks", data)
}

// Handler for submitting a new task.
func (h TaskHandler) FormSubmitHandler(context echo.Context) error {
	newTask := usecases.NewTask{
		Title:       context.FormValue("title"),
		Description: context.FormValue("description"),
		UserId:      context.FormValue("user-id"),
		CategoryId:  context.FormValue("category-id"),
		StatusId:    context.FormValue("status-id"),
	}
	_, err := h.TaskUsecase.Create(&newTask)
	if err != nil {
		return err
	}
	tasks, err := h.TaskUsecase.GetAll()
	if err != nil {
		return err
	}
	return context.Render(http.StatusOK, "tasksList", tasks)
}

// Handler for listing all tasks.
func (h TaskHandler) ListHandler(context echo.Context) error {
	tasks, err := h.TaskUsecase.GetAll()
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	return context.Render(http.StatusOK, "tasksList", tasks)
}

// Handler for deleting a task.
func (h TaskHandler) DeleteHandler(context echo.Context) error {
    id := context.Param("id")
    err := h.TaskUsecase.Delete(id)
    if err != nil {
        return err
    }

    tasks, err := h.TaskUsecase.GetAll()
    if err != nil {
        return err
    }

    return context.Render(http.StatusOK, "tasksList", tasks)
}

// Handler for getting a single task.
func (h TaskHandler) GetTaskHandler(context echo.Context) error {
	id := context.Param("id")
	task, err := h.TaskUsecase.GetById(id)
	if err != nil {
		return err
	}

	return context.Render(http.StatusOK, "taskDetail", task)
}
