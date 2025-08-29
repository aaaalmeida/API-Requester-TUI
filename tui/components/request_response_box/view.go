package request_response_box

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(65).
		Height(25).
		Render("response box")

	return bodyBox
}
