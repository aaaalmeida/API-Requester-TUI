package main_page

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {

	leftSide := lipgloss.JoinVertical(
		lipgloss.Top,
		m.subcomponents[SearchCollectionIndex].View(),
		m.subcomponents[CollectionMenuIndex].View(),
	)

	rightSide := lipgloss.JoinVertical(
		lipgloss.Top,
		m.subcomponents[HeaderIndex].View(),
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.subcomponents[RequestHeadersIndex].View(),
			m.subcomponents[RequestResponseIndex].View()),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftSide, rightSide)
}
