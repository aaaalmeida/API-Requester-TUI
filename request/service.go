package request

import (
	"api-requester/context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

/**
* 	Create and insert request into database. Status_code, headers and body are optional and can
*	be pass as nil. Created_at and updated_at are automatically set as local date time.
*	Returns pointer to created request.
 */
func AddRequest(ctx *context.AppContext, name, url string,
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

/*
*	Return array of all requests.
 */
func GetAllRequest(ctx *context.AppContext) ([]Request, error) {
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

/*
*	Return array of all requests with matching method_id.
 */
func SearchRequestByMethodId(ctx *context.AppContext, method_id int) ([]Request, error) {
	rows, err := ctx.DB.Query("SELECT * FROM request WHERE method_id = ?;", method_id)
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
		)
		if err != nil {
			return nil, err
		}

		requests = append(requests, req)
	}
	return requests, nil
}

/*
*	Return request with matching id or ErrNoRows if not found.
 */
func SearchRequestById(ctx *context.AppContext, request_id int) (*Request, error) {
	row := ctx.DB.QueryRow(`SELECT id, name, url, method_id, collection_id, status_code, headers, body
	 FROM request WHERE id = ?;`, request_id)

	var request Request
	err := row.Scan(
		&request.ID,
		&request.Name,
		&request.Url,
		&request.Method_id,
		&request.Collection_id,
		&request.Status_code,
		&request.Headers,
		&request.Body,
	)

	if err != nil {
		return nil, err
	}
	return &request, nil
}

/*
*	Return array of all requests with matching collection_id.
 */
func SearchRequestByCollectionId(ctx *context.AppContext, collection_id int) ([]Request, error) {
	rows, err := ctx.DB.Query("SELECT * FROM request WHERE collection_id = ?;", collection_id)
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

// TODO: remover referencia do ponteiro da lista de requests durante execução
/*
*	Delete request with matching id from database.
 */
func DeleteRequestById(ctx *context.AppContext, request_id int) error {
	stmt, err := ctx.DB.Prepare("DELETE FROM request WHERE id = ?;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request_id)
	return err
}

/*
*	Update and saves request in db. This DOES NOT override uninformed values, only replaces new ones.
*	ID, Collection_ID, Created_at and Updated_at will be ignored because this function is not mean
*	to update then manually. Updated_at is automatic updated to local date time.
 */
func UpdateRequest(ctx *context.AppContext, request_id int, request *Request) error {
	queryClauses := []string{}
	args := []interface{}{}

	if request.Name != "" {
		queryClauses = append(queryClauses, "name = ?")
		args = append(args, request.Name)
	}

	if request.Url != "" {
		queryClauses = append(queryClauses, "url = ?")
		args = append(args, request.Url)
	}

	if request.Method_id != 0 {
		queryClauses = append(queryClauses, "method_id = ?")
		args = append(args, request.Method_id)
	}

	if request.Status_code.Valid {
		queryClauses = append(queryClauses, "status_code = ?")
		args = append(args, request.Status_code.Int16)
	}

	if request.Headers.Valid {
		queryClauses = append(queryClauses, "headers = ?")
		args = append(args, request.Headers.String)
	}

	if request.Body.Valid {
		queryClauses = append(queryClauses, "body = ?")
		args = append(args, request.Body.String)
	}

	if len(queryClauses) == 0 {
		return fmt.Errorf("nothing to update")
	}

	queryClauses = append(queryClauses, "updated_at = ?")
	args = append(args, time.Now().Format(time.DateTime))

	query := fmt.Sprintf("UPDATE request SET %s WHERE id = ?;", strings.Join(queryClauses, ", "))
	args = append(args, request_id)
	_, err := ctx.DB.Exec(query, args...)
	return err
}

/*
*	Call HTTP Request using Headers (if valid). Returns body stringified.
 */
func CallRequest(req *Request) (string, error) {
	// create request body if it is valid
	var body = strings.NewReader("")
	if req.Body.Valid {
		body = strings.NewReader(req.Body.String)
	}

	method := http.MethodGet
	switch req.Method_id {
	case 2:
		method = http.MethodPost
	case 3:
		method = http.MethodPut
	case 4:
		method = http.MethodDelete
	case 5:
		method = http.MethodPatch
	case 6:
		method = http.MethodHead
	case 7:
		method = http.MethodTrace
	}

	httpRequest, err := http.NewRequest(method, req.Url, body)
	if err != nil {
		return "", err
	}

	// TODO: depois que implementar a TUI, remover esse header
	if req.Body.Valid {
		httpRequest.Header.Set("Content-type", "application/json")
	}

	// set request headers if valid
	if req.Headers.Valid {
		var headersMap map[string]string
		if err := json.Unmarshal([]byte(req.Headers.String), &headersMap); err != nil {
			return "", err
		}

		for key, value := range headersMap {
			httpRequest.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

/*
*	Call HTTP Request with matching id. Returns body stringified.
 */
func CallRequestById(ctx *context.AppContext, request_id int) (string, error) {
	row := ctx.DB.QueryRow("SELECT * FROM request WHERE id = ?;", request_id)

	var request Request
	err := row.Scan(
		&request.ID,
		&request.Url,
		&request.Name,
		&request.Method_id,
		&request.Collection_id,
		&request.Status_code,
		&request.Headers,
		&request.Body,
		&request.Created_at,
		&request.Updated_at)

	if err != nil {
		return "", err
	}

	return CallRequest(&request)
}
