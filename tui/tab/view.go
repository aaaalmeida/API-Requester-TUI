package tab

import (
	"api-requester/utils"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	var renderedTabs []string

	for i, tab := range m.tabs {
		isFirst := i == 0
		isLast := i == len(m.tabs)-1
		isActive := i == m.activeTab

		style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)

		style = style.Inherit(isFirstStyle(isFirst))
		style = style.Inherit(isLastStyle(isLast))
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
		return lipgloss.NewStyle().BorderBottom(false)
	}
	return lipgloss.NewStyle()
}
