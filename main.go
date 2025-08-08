package main

import (
	"database/sql"
	"fmt"
	"log"

	"api-requester/appctx"
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

	requests, err := request.GetAllRequest(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range requests {
		fmt.Println(i.Name, i.Url, i.Headers, i.Body, i.Method_id)
	}
}
