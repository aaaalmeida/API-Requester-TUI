package tab_body

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	selectMethod := m.selectMethod.View()
	inputUrl := m.input.View()

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(60).
		Height(25).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, selectMethod, inputUrl))

	return bodyBox
}
