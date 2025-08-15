package collection_menu

import (
	"api-requester/context"
	"api-requester/request"

	tea "github.com/charmbracelet/bubbletea"
)

func fetchRequestsFromCollectionCmd(ctx *context.AppContext, collection_id int) tea.Cmd {
	return func() tea.Msg {
		reqs, err := request.SearchRequestByCollectionId(ctx, collection_id)
		return requestLoadedMsg{
			collection_id: collection_id,
			requests:      reqs,
			err:           err,
		}
	}
}
