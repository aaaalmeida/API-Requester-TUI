package main

import (
	"api-requester/context"
	"api-requester/tui/collection"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, err := context.NewAppContext()
	if err != nil {
		log.Fatalln(err)
	}

	p := tea.NewProgram(collection.NewModel(ctx))
	_, err = p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
