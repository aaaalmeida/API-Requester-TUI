package button

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	normalStyle := lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.NormalBorder())
	focusedStyle := normalStyle.Inherit(lipgloss.NewStyle().Background(lipgloss.Color("2")))

	if !m.isFocused {
		return normalStyle.Render(m.text)
	} else {
		return focusedStyle.Render(m.text)
	}
}
