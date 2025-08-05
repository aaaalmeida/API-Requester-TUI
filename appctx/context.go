package appctx

import (
	"database/sql"
)

// GLOBAL CONTEXT REFERENCE
type AppContext struct {
	DB *sql.DB
}

func NewAppContext(db *sql.DB) *AppContext {
	return &AppContext{
		DB: db,
	}
}
