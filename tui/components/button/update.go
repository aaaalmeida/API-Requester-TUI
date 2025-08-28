package button

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg.(tea.Key).String() == "enter" {
		return m, m.onPress
	}

	return m, nil
}
