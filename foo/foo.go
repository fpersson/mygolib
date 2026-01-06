package foo

import (
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

// Pointless
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "notworking")
	if err != nil {
		slog.Error(err.Error())
	}

	return db, nil
}
