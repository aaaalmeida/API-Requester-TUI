package tab

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var view strings.Builder

	tab_header := m.tab_header.View()

	// BODY
	tab_body := m.tab_body.View()

	// JOIN HEADER AND BODY
	view.WriteString(lipgloss.JoinVertical(lipgloss.Top, tab_header, tab_body))

	return view.String()
}
