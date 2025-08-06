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
		var collec Collection
		// copy row into collec variable
		err := rows.Scan(&collec.ID, &collec.Name, &collec.Created_at, &collec.Description)
		if err != nil {
			return nil, err
		}

		// add collec into collections
		collections = append(collections, collec)
	}

	return collections, nil
}

func AddCollection(ctx *appctx.AppContext, name string, description *string) (*Collection, error) {
	// prepared statment, secure against sql injection
	stmt, err := ctx.DB.Prepare("INSERT INTO collection (name, description) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// go doesnt allow strings to be null, so I must use a pointer to string
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
		ID:          int(lastId),
		Name:        name,
		Created_at:  time.Now().String(),
		Description: desc,
	}, nil
}
