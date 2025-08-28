package request_header_box

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	selectMethod := m.selectMethod.View()
	inputUrl := m.inputMethod.View()
	sendRequestButton := m.sendRequestButton.View()

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(65).
		Height(25).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, selectMethod, inputUrl, sendRequestButton))

	return bodyBox
}
