package main

import (
	"database/sql"
	"fmt"
	"log"

	"api-requester/appctx"
	"api-requester/collection"
	"api-requester/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConnection, err := sql.Open("sqlite3", "mydb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	db.InitSchema(dbConnection, "db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Created")

	// global context
	ctx := appctx.NewAppContext(dbConnection)

	// collection.AddCollection(ctx, "teste1", nil)
	// collection.AddCollection(ctx, "teste2", nil)
	// collection.AddCollection(ctx, "teste3", nil)
	// collection.AddCollection(ctx, "teste4", nil)
	// desc := "qweqwe"
	// collection.AddCollection(ctx, "teste5", &desc)
	// collection.AddCollection(ctx, "teste6", &desc)
	// collection.AddCollection(ctx, "teste7", &desc)
	// collection.AddCollection(ctx, "teste8", &desc)

	colls, err := collection.GetAllCollection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range colls {
		fmt.Println(i)
	}
}
