package tab

import (
	"api-requester/utils"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	var renderedTabs []string

	for i, tab := range m.tabs {
		isActive := i == m.activeTab

		style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
		style = style.Inherit(isActiveStyle(isActive))

		renderedTabs = append(renderedTabs, style.Render(tab.title))
	}

	var doc strings.Builder

	tabsHeader := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	body := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(60).
		Height(22).
		Render(utils.LoremIpsup())

	doc.WriteString(lipgloss.JoinVertical(lipgloss.Top, tabsHeader, body))
	return doc.String()
}

func isActiveStyle(value bool) lipgloss.Style {
	if value {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("2"))
	}
	return lipgloss.NewStyle()
}
