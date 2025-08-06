package request

import (
	"database/sql"
)

type Request struct {
	ID            int
	Name          string
	Url           string
	Method_id     int
	Collection_id int
	Status_code   sql.NullInt16
	Headers       sql.NullString
	Body          sql.NullString
	Created_at    string
	Updated_at    string
}
