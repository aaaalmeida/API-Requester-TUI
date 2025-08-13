package collection

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m CollectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// captures msg type
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// captures which key was pressed
		switch msg.String() {

		// exit program
		case "ctrl+q", "esc":
			return m, tea.Quit

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Collections)-1 {
				m.Cursor++
			}
		}
	}

	return m, nil
}
