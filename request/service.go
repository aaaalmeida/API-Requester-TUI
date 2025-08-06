package request

import (
	"api-requester/appctx"
	"database/sql"
	"fmt"
	"time"
	// "bytes"
	// "fmt"
	// "net/http"
	// "strings"
)

func AddRequest(ctx *appctx.AppContext, name, url string,
	method_id, collection_id int,
	status_codePointer *int,
	headersPointer, bodyPointer *string) (*Request, error) {

	stmt, err := ctx.DB.Prepare(`
		INSERT INTO request(name, url, method_id, collection_id, status_code, headers, body)
		 VALUES (?, ?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var status_code sql.NullInt16
	if status_codePointer != nil {
		status_code = sql.NullInt16{Valid: true, Int16: int16(*status_codePointer)}
	} else {
		status_code = sql.NullInt16{Valid: false}
	}

	var headers, body sql.NullString
	if headersPointer != nil {
		headers = sql.NullString{Valid: true, String: *headersPointer}
	} else {
		headers = sql.NullString{Valid: false}
	}
	if bodyPointer != nil {
		body = sql.NullString{Valid: true, String: *bodyPointer}
	} else {
		body = sql.NullString{Valid: false}
	}

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(name, url, method_id, collection_id, status_code, headers, body)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Request{
		ID:            int(lastId),
		Name:          name,
		Url:           url,
		Method_id:     method_id,
		Collection_id: collection_id,
		Status_code:   status_code,
		Headers:       headers,
		Body:          body,
		Created_at:    time.Now().String(),
		Updated_at:    time.Now().String(),
	}, nil
}

func GetAllRequest(ctx *appctx.AppContext) ([]Request, error) {
	// FIXME: possivel erro qnd query puxa todos dados
	rows, err := ctx.DB.Query("SELECT * FROM request;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(
			&req.ID,
			&req.Name,
			&req.Url,
			&req.Method_id,
			&req.Collection_id,
			&req.Status_code,
			&req.Headers,
			&req.Body,
			&req.Created_at,
			&req.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func SearchRequestByMethodId(ctx *appctx.AppContext, method_id int) ([]Request, error) {
	rows, err := ctx.DB.Query(`
		SELECT name, url, method_id, collection_id, status_code, headers, body
		FROM request WHERE method_id = ?;`,
		method_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(&req.ID,
			&req.Name,
			&req.Url,
			&req.Method_id,
			&req.Collection_id,
			&req.Status_code,
			&req.Headers,
			&req.Body)
		if err != nil {
			return nil, err
		}

		requests = append(requests, req)
	}
	return requests, nil
}

// TODO:
// func UpdateRequest(ctx *appctx.AppContext){}

func CallRequest(req *Request) {
	// create request body if it is valid
	// var body *strings.Reader
	// if req.Body.Valid {
	// 	body = strings.NewReader(req.Body.String)
	// }else {
	// 	body = strings.NewReader("")
	// }

	// var method string
	// switch req.Method_id {
	// 	case 1:
	// 		method = http.MethodGet
	// 	case 2:
	// 		method = http.MethodPost
	// 	case 3:
	// 		method = http.MethodPut
	// 	case 4:
	// 		method = http.MethodDelete
	// 	case 5:
	// 		method = http.MethodPatch
	// 	case 6:
	// 		method = http.MethodHead
	// 	case 7:
	// 		method = http.MethodTrace
	// }

	// httpRequest, err := http.NewRequest(method, req.Url, body)
	// if err != nil {
	// 	return "", err
	// }

	// httpRequest.Header.Set()

	fmt.Printf("name %s\n", req.Name)
	fmt.Printf("url %s\n", req.Url)
	fmt.Printf("status %d\n", req.Status_code.Int16)
	fmt.Printf("header %s\n", req.Headers.String)
	fmt.Printf("body %s\n\n", req.Body.String)
}
