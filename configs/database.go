package configs

import (
	"database/sql"
	"log"
	"os"
)

type DB interface {
	Exec(query string, args ...any) (sql.Result, error)
}

type SQL interface {
	Open(driverName, dataSourceName string) (*DB, error)
}

func AutoMigrate(db DB) {
	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err := db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
}

func ConnectDatabase() *sql.DB {
	var err error

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Connect database error", err)
	}

	return db
}
