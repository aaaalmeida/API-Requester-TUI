package db

import (
	"database/sql"
	"os"
	"strings"
)

func InitSchema(db *sql.DB, path string) error {
	// read schema.sql
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// split sql into commands
	commands := strings.Split(string(content), ";")

	// execute each command
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		_, err := db.Exec(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
