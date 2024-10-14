package repository

import (
	"database/sql"
	"log/slog"
	"os"
	_ "github.com/lib/pq"
)

func ConnectDatabase(logger *slog.Logger) (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed connect database")
		return nil, err
	}
	logger.Debug("Connect database")
	return db, nil
}