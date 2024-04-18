package store

import (
	"database/sql"
	"os"
	"strings"
	"web-forum/configs"

	_ "github.com/mattn/go-sqlite3" // go get github.com/mattn/go-sqlite3
)

func NewSqlite3(config configs.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.Driver, config.DB.Dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = CreateTables(db, config); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB, config configs.Config) error {
	file, err := os.ReadFile(config.Migrate)
	if err != nil {
		return err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}
	return err
}
