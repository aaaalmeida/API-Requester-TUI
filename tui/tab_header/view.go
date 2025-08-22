package tab_header

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	style = style.Inherit(isCursorPointingStyle(m.isSelected))

	return style.Render(m.request.Name)
}

func isCursorPointingStyle(value bool) lipgloss.Style {
	if value {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("2"))
	}
	return lipgloss.NewStyle()
}
