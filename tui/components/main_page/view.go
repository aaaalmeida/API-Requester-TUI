package main_page

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	leftSide := lipgloss.JoinVertical(
		lipgloss.Top,
		m.subcomponents[0].View(),
		m.subcomponents[1].View(),
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftSide,
		m.subcomponents[2].View(),
	)
}
