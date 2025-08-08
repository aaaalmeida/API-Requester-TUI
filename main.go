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

	req, err := request.SearchRequestById(ctx, 3)
	if err != nil {
		log.Fatal(err)
	}
	request.CallRequest(req)
	// reqs, err := request.GetAllRequest(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i, req := range reqs {
	// 	response, err := request.CallRequest(&req)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(i, response)
	// }
}
