package repository

import (
	_"github.com/lib/pq"
	"database/sql"
	// "os"
	// "log/slog"
)

var database *sql.DB

// func ConnectDatabase() {
// 	connStr := os.Getenv("DATABASE_URL")
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		logger.Error("Failed connect database")
//         panic(err)
//     } 
// 	logger.Debug("Connect database")
// 	database = db
// }