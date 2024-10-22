package main

import (
	repositorymysql "taskmanager/src/adapters/repository_mysql"
	"taskmanager/src/api"
	"taskmanager/src/usecases"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pool := repositorymysql.GetDBPool()

	// define usecases
	TaskUsecase := usecases.Task{
		Repo: repositorymysql.Task{Pool: pool},
	}
	CategoryUsecase := usecases.Category{
		Repo: repositorymysql.Category{Pool: pool},
	}
	StatusUsecase := usecases.Status{
		Repo: repositorymysql.Status{Pool: pool},
	}
	UserUsecase := usecases.User{
		Repo: repositorymysql.User{Pool: pool},
	}

	// define api server
	api := api.API{
		UseCaseTask:     TaskUsecase,
		UseCaseCategory: CategoryUsecase,
		UseCaseStatus:   StatusUsecase,
		UseCaseUser:     UserUsecase,
	}

	// run server
	api.Run()

	defer pool.Close()
}
