package request_header_box

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	// selectMethod := m.selectMethod.View()
	// inputUrl := m.inputUrl.View()
	// buttonSendRequest := m.buttonSendRequest.View()

	var views []string
	for _, sc := range m.subcomponents {
		views = append(views, sc.View())
	}

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(65).
		Height(25).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, views...))

	return bodyBox
}
