package repository_mysql

import (
	"database/sql"
)

// GetDBPool returns a pool of connections to the database
func GetDBPool() *sql.DB {
	pool, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/task_manager")
	if err != nil {
		panic(err)
	}
	return pool
}
