package collection

import (
	"api-requester/appctx"
	"database/sql"
	"time"
)

func GetAllCollection(ctx *appctx.AppContext) ([]Collection, error) {
	rows, err := ctx.DB.Query("SELECT id, name, created_at, description FROM collection")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var collections []Collection

	for rows.Next() {
		var c Collection
		// copy row into c variable
		err := rows.Scan(&c.id, &c.name, &c.created_at, &c.description)
		if err != nil {
			return nil, err
		}

		// add c into collections
		collections = append(collections, c)
	}

	return collections, nil
}

func AddCollection(ctx *appctx.AppContext, name string, description *string) (*Collection, error) {
	// prepared statment
	// secure against sql injection
	stmt, err := ctx.DB.Prepare("INSERT INTO collection (name, description) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// description can be null
	var desc sql.NullString
	if description != nil {
		desc = sql.NullString{String: *description, Valid: true}
	} else {
		desc = sql.NullString{Valid: false}
	}

	result, err := stmt.Exec(name, desc)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Collection{
		id:          int(lastId),
		name:        name,
		created_at:  time.Now().String(),
		description: desc,
	}, nil
}
