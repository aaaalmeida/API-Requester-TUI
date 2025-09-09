package select_menu

import (
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.isOpened = !m.isOpened
		case "esc":
			m.isOpened = false
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
		m.ctx.Logger.Println("OQ VEIO NO CMD", msg)

		if msg.Err != nil {
			//TODO: tratar erro
		}
		// only load methods once
		if m.Options == nil {
			m.Options = ConvertMethodsToSelectOptions(msg.Methods)
			m.ctx.Logger.Println("OQ SETOU NO MODEL", msg)
		}
	}

	return m, nil
}
