package select_menu

import (
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if m.isOpened {
				m.selectedItem = m.cursor
				m.isOpened = false
			} else {
				m.isOpened = true
			}

		case "a":
			m.ctx.Logger.Println("teste")
		case "up":
			m.ctx.Logger.Println("cima")
			if m.isOpened && m.cursor > 0 {
				m.cursor--
			}
		case "down":
			m.ctx.Logger.Println("baixo")
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
	}

	return m, nil
}
