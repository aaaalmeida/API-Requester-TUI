package button

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	return style.Inherit(isFocusedStyle(m.Focused)).Render(m.label)
}

func isFocusedStyle(b bool) lipgloss.Style {
	if b {
		return lipgloss.NewStyle().Background(lipgloss.Color("4"))
	}
	return lipgloss.NewStyle()
}
