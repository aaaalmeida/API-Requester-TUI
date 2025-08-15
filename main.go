package main

import (
	"api-requester/context"
	"api-requester/tui/main_page"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, err := context.NewAppContext()
	if err != nil {
		log.Fatalln(err)
	}

	p := tea.NewProgram(main_page.NewModel(ctx))
	_, err = p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
