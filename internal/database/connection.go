package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const m_DRIVER_NAME = "mysql"

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(m_DRIVER_NAME, "test:test@/project-calculator")
	if err != nil {
		log.Fatal("Error opening database:", err)
		return nil, err
	}
	defer db.Close()
	return db, nil
}
