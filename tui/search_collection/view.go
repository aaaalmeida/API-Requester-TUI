package search_collection

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(30).
		Render(m.textInput.View())
}
