package tab

import (
	"api-requester/utils"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var view strings.Builder
	var renderedTabs []string

	// HEADER
	for _, req := range m.tabsHeader {
		renderedTabs = append(renderedTabs, req.View())
	}

	// JOIN TABS INSIDE HEADER BOX
	tabsContent := lipgloss.JoinHorizontal(lipgloss.Left, renderedTabs...)
	tabsHeaderBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(130).
		Height(3).
		Render(tabsContent)

	// BODY
	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		Width(60).
		Height(25).
		// TODO: TIRAR LOREM IPSUM
		Render(utils.LoremIpsum())

	// JOIN HEADER AND BODY
	view.WriteString(lipgloss.JoinVertical(lipgloss.Top, tabsHeaderBox, bodyBox))

	return view.String()
}
