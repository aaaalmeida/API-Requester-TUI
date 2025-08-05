package collection

import "database/sql"

type Collection struct {
	id          int
	name        string
	created_at  string
	description sql.NullString
}
