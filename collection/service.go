package collection

import (
	"api-requester/appctx"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func GetAllCollection(ctx *appctx.AppContext) ([]Collection, error) {
	rows, err := ctx.DB.Query("SELECT id, name, created_at, updated_at, description FROM collection")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var collections []Collection

	for rows.Next() {
		var collec Collection
		err := rows.Scan(&collec.ID, &collec.Name, &collec.Created_at, &collec.Updated_at, &collec.Description)
		if err != nil {
			return nil, err
		}

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

	// go doesnt allow strings to be null, so I must use sql nullString
	var desc sql.NullString
	if description != nil {
		desc = sql.NullString{String: *description, Valid: true}
	} else {
		desc = sql.NullString{Valid: false}
	}

	result, err := stmt.Exec(name, desc.String)
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

func UpdateCollection(ctx *appctx.AppContext, collection_id int, collection *Collection) error {
	queryClauses := []string{}
	args := []interface{}{}

	if collection.Name != "" {
		queryClauses = append(queryClauses, "name = ?")
		args = append(args, collection.Name)
	}

	if collection.Description.Valid {
		queryClauses = append(queryClauses, "description = ?")
		args = append(args, collection.Description.String)
	}

	if len(queryClauses) == 0 {
		return fmt.Errorf("nothing to update")
	}

	queryClauses = append(queryClauses, "updated_at = ?")
	args = append(args, time.Now().Format(time.DateTime))

	query := fmt.Sprintf("UPDATE collection SET %s WHERE id = ?;", strings.Join(queryClauses, ", "))
	args = append(args, collection_id)
	_, err := ctx.DB.Exec(query, args...)
	return err
}
