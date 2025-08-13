package context

import (
	"api-requester/db"
	"database/sql"
)

// GLOBAL CONTEXT REFERENCE
type AppContext struct {
	DB *sql.DB
}

func NewAppContext() (*AppContext, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	err = db.InitSchema(database)
	if err != nil {
		return nil, err
	}

	return &AppContext{
		DB: database,
	}, nil
}
