package request_header_box

import (
	cmds "api-requester/tui/commands"
	"api-requester/tui/components/select_menu"
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "left":

		case "right":
		case "enter":
		}

	// CAPTURE REQUEST
	case messages.LoadRequestMsg:
		m.request = msg.Request

		return m, cmds.FetchMethodsCmd(m.context)

	// FETCHES METHODS FROM DB AND SEND TO SELECT COMPONENT
	case messages.LoadMethodsMsg:
		m.selectMethod.Options = select_menu.ConvertMethodsToSelectOptions(msg.Methods)
	}

	return m, nil
}
