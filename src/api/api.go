package api

import (
	"html/template"
	"io"
	"path"
	"taskmanager/src/api/handlers"
	"taskmanager/src/usecases"

	"github.com/labstack/echo/v4"
)

type API struct {
	UseCaseTask     usecases.ITask
	UseCaseCategory usecases.ICategory
	UseCaseStatus   usecases.IStatus
	UseCaseUser     usecases.IUser
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (api API) Run() {

	taskHandler := handlers.TaskHandler{
		TaskUsecase:     api.UseCaseTask,
		CategoryUsecase: api.UseCaseCategory,
		StatusUsecase:   api.UseCaseStatus,
		UserUsecase:     api.UseCaseUser,
	}
	userHandler := handlers.UserHandler{
		UserUsecase: api.UseCaseUser,
	}

	// Create a new Echo instances
	e := echo.New()

	// add templates
	var templateRoot = "src/web/templates"
	t := &Template{
		templates: template.Must(template.ParseGlob(path.Join(templateRoot, "*", "*.html"))),
	}
	e.Renderer = t

	// Middleware
	// e.Use(middleware.Logger())

	// Serve static files
	e.Static("/static/", path.Join("src", "web", "styles", "static"))

	// pages
	e.GET("/", handlers.Root)
	e.GET("/tasks", taskHandler.BaseHandler)
	e.GET("/users", userHandler.BaseHandler)

	// dark mode
	e.POST("/toggle-dark-mode", handlers.ToggleDarkMode)

	// components - tasks
	e.POST("/content/tasks-form", taskHandler.FormSubmitHandler)
	e.GET("/content/tasks-list", taskHandler.ListHandler)
	e.DELETE("/content/tasks/:id", taskHandler.DeleteHandler)
	e.GET("/content/tasks/:id", taskHandler.GetTaskHandler)

	// components - users
	e.GET("content/users", userHandler.ListHandler)
	e.POST("content/users-form", userHandler.FormSubmitHandler)
	e.DELETE("content/users/:id", userHandler.DeleteHandler)


	// Start server
	e.Logger.Fatal(e.Start(":7070"))
}
