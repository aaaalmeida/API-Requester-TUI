package main_page

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	var subComps []string

	for _, sc := range m.subcomponents {
		subComps = append(subComps, sc.View())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, subComps...)
}
