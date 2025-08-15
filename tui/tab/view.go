package tab

import (
	"api-requester/utils"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	defaultBoxTitleStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	var renderedTabs []string

	for i, tab := range m.tabs {
		isFirst := i == 0
		isLast := i == len(m.tabs)-1
		isActive := i == m.activeTab

		defaultBoxTitleStyle.
			Inherit(isFirstStyle(isFirst)).
			Inherit(isLastStyle(isLast)).
			Inherit(isActiveStyle(isActive))

		renderedTabs = append(renderedTabs, defaultBoxTitleStyle.Render(tab.title))
	}

	var doc strings.Builder

	tabsHeader := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	body := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(150).
		Height(20).
		Render(utils.LoremIpsup())

	doc.WriteString(lipgloss.JoinVertical(lipgloss.Left, tabsHeader, body))
	return doc.String()
}

func isFirstStyle(value bool) lipgloss.Style {
	if value {
		return lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, false)
	}
	return lipgloss.NewStyle()
}

func isLastStyle(value bool) lipgloss.Style {
	if value {
		return lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, false, true, true)
	}
	return lipgloss.NewStyle()
}

func isActiveStyle(value bool) lipgloss.Style {
	if value {
		return lipgloss.NewStyle().BorderBottom(false).Background(lipgloss.Color("4"))
	}
	return lipgloss.NewStyle()
}
