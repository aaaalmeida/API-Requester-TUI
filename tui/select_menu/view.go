package select_menu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	// TODO: melhorar interface
	selectedStyle := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("2")).Border(lipgloss.NormalBorder())
	normalStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	if len(m.options) == 0 {
		return selectedStyle.Render("")
	}

	if !m.isOpened {
		return selectedStyle.Render(m.options[m.selectedItem].Label())
	}

	var b strings.Builder
	for i, opt := range m.options {
		if i == m.cursor {
			b.WriteString(selectedStyle.Render(opt.Label()) + "\n")
		} else {
			b.WriteString(normalStyle.Render(opt.Label()) + "\n")
		}
	}

	return b.String()
}
