package db

import (
	"database/sql"
	"os"
)

func InitSchema(db *sql.DB, path string) error {
	// read schema.sql
	content, err := os.ReadFile(path)
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
