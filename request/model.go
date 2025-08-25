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

// func (r Request) String() string {
// 	return fmt.Sprintf("%p %d %s %s %d %d %d %s %s %s %s", &r, r.ID, r.Name, r.Url, r.Method_id, r.Collection_id, r.Status_code.Int16, r.Headers.String, r.Body.String, r.Created_at, r.Updated_at)
// }

func (r Request) String() string {
	return fmt.Sprintf("%p", &r)
}
