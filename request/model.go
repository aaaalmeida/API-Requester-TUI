package request

import (
	"database/sql"
	"fmt"
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

func (r Request) String() string {
	return fmt.Sprintf("%d %s %s %d", r.ID, r.Name, r.Url, r.Collection_id)
}
