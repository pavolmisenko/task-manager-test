package handlers

import (
	"net/http"
	"taskmanager/src/usecases"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecases.IUser
}

func (u UserHandler) BaseHandler(context echo.Context) error {
	darkMode := false
    if cookie, err := context.Cookie("dark-mode"); err == nil {
        darkMode = cookie.Value == "true"
    }

	data := struct {
		DarkMode bool
	}{
		DarkMode: darkMode,
	}
	return context.Render(http.StatusOK, "users", data)
}

func (u UserHandler) FormSubmitHandler(context echo.Context) error {
	newUser := usecases.NewUser{
		Name:  context.FormValue("name"),
		Email: context.FormValue("email"),
	}
	_, err := u.UserUsecase.Create(&newUser)
	if err != nil {
		return err
	}

	users, err := u.UserUsecase.GetAll()
	if err != nil {
		return err
	}

	return context.Render(http.StatusOK, "usersList", users)
}

func (u UserHandler) ListHandler(context echo.Context) error {
	users, err := u.UserUsecase.GetAll()

	if err != nil {
		return err
	}
	return context.Render(http.StatusOK, "usersList", users)
}

func (u UserHandler) DeleteHandler(context echo.Context) error {
	id := context.Param("id")
	err := u.UserUsecase.Delete(id)
	if err != nil {
		return err
	}

	users, err := u.UserUsecase.GetAll()
	if err != nil {
		return err
	}

	return context.Render(http.StatusOK, "usersList", users)
}
