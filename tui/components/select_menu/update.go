package select_menu

import (
	"api-requester/tui/commands"
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if m.isOpened {
				// close menu
				m.selectedItem = m.cursor
				m.isOpened = false
				return m, commands.UserPressEnterInSelect(m.Options[m.selectedItem])
			} else {
				// open menu
				m.isOpened = true
			}

		case "up":
			if m.isOpened && m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.isOpened && m.cursor < len(m.Options)-1 {
				m.cursor++
			}
		}

		// INITIAL UPDATE
	case messages.LoadMethodsMsg:
		if msg.Err != nil {
			//TODO: tratar erro
		}
		// only load methods once
		if m.Options == nil {
			m.Options = ConvertMethodsToSelectOptions(msg.Methods)
		}

	case messages.SendNumberMsg:
		if msg.Err != nil {
			//TODO: tratar erro
			m.ctx.Logger.Println(msg.Err)
		}

		// sqlite index start at 1, so we subtract from index
		m.selectedItem = msg.Value - 1
	}

	return m, nil
}
