package collection

import "database/sql"

type Collection struct {
	ID          int
	Name        string
	Created_at  string
	Updated_at  string
	Description sql.NullString
}
