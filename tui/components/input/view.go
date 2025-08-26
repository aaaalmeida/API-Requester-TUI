package input

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(30).
		Height(3).
		Render(m.textInput.View())
}
