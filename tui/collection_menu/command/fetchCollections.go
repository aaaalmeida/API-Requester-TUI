package collection_menu

import (
	"api-requester/collection"
	"api-requester/context"
	msg "api-requester/tui/collection_menu/msg"

	tea "github.com/charmbracelet/bubbletea"
)

func FetchCollectionsCmd(ctx *context.AppContext) tea.Cmd {
	return func() tea.Msg {
		cols, err := collection.GetAllCollection(ctx)

		return msg.LoadCollectionMsg{
			Collections: cols,
			Err:         err,
		}
	}
}
