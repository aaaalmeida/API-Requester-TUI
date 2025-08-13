package db

import (
	"database/sql"
	"os"
)

const DB_NAME string = "mydb.sqlite"
const SCHEMA string = "db/schema.sql"

func InitSchema(db *sql.DB) error {
	// read schema.sql
	content, err := os.ReadFile(SCHEMA)
	if err != nil {
		return err
	}

	// sqlite separates instructions alone
	sql := string(content)
	_, err = db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", DB_NAME)
}
