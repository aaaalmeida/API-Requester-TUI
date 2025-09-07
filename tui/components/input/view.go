package input

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	normalStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	focusedStyle := normalStyle.Inherit(lipgloss.NewStyle().Background(lipgloss.Color("2")))

	if !m.isFocused {
		return normalStyle.Render(m.textInput.View())
	} else {
		return focusedStyle.Render(m.textInput.View())
	}
}
