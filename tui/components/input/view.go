package input

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Render(m.textInput.View())
}
