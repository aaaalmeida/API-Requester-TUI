package input

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	if m.isFocused {
		return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).
			Background(lipgloss.Color("2")).
			Render(m.textInput.View())
	} else {
		return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(m.textInput.View())
	}
}
