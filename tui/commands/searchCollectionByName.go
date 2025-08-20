package commands

import (
	"api-requester/collection"
	"api-requester/context"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func SearchCollectionByNameCmd(ctx *context.AppContext, collection_name string) tea.Cmd {
	return func() tea.Msg {
		cols, err := collection.SearchCollectionByName(ctx, collection_name)

		return msg.LoadCollectionsMsg{
			Collections: cols,
			Err:         err,
		}
	}
}
