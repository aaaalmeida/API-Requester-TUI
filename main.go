package main

import (
	"database/sql"
	"fmt"
	"log"

	"api-requester/appctx"
	"api-requester/collection"
	"api-requester/db"
	"api-requester/request"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConnection, err := sql.Open("sqlite3", "mydb.sqlite")
	if err != nil {
		log.Fatal("error opening db ", err)
	}
	defer dbConnection.Close()

	err = db.InitSchema(dbConnection, "db/schema.sql")
	if err != nil {
		log.Fatal("error creating db ", err)
	}
	fmt.Println("DB Created")

	// global context
	ctx := appctx.NewAppContext(dbConnection)

	collection.AddCollection(ctx, "col1", nil)
	collection.AddCollection(ctx, "col2", nil)
	collection.AddCollection(ctx, "col3", nil)
	var desc = "eqweqweqweqwe"
	collection.AddCollection(ctx, "col4", &desc)
	collection.AddCollection(ctx, "col5", &desc)
	collection.AddCollection(ctx, "col6", &desc)
	collection.AddCollection(ctx, "col7", &desc)

	// var status = 200
	// var header = "{\n  'hello': 'world',\n  \"book\": \"on the table\",\n  \"happy\" : false,\n  \"age\": 24,\n  [\n    \"eyes\": \"brown\",\n    \"hair\": \"black\",\n    \"favorite_color\": \"blue\"\n  ]\n}"
	// var body = "{\n  [\n    \"eyes\": \"brown\",\n    \"hair\": \"black\",\n    \"favorite_color\": \"blue\"\n  ]\n}"

	// req3, _ := request.AddRequest(ctx, "req3", "https://example.com", 1, 1, &status, &header, &body)
	// request.CallRequest(req3)

	reqs, _ := request.GetAllRequest(ctx)

	for i, r := range reqs {
		fmt.Printf("%d--------------------\n", i)
		fmt.Printf("name %s\n", r.Name)
		fmt.Printf("url %s\n", r.Url)
		fmt.Printf("status %d\n", r.Status_code.Int16)
		fmt.Printf("header %s\n", r.Headers.String)
		fmt.Printf("body %s\n\n", r.Body.String)
	}
}
