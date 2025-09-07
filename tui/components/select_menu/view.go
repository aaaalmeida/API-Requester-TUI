package select_menu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	normalStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(m.width)
	selectedStyle := normalStyle.Inherit(lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("2")))

	if len(m.Options) == 0 {
		return normalStyle.Render("---")
	}

	if !m.isOpened {
		return normalStyle.Render(m.Options[m.selectedItem].Label())
	}

	var b strings.Builder
	for i, opt := range m.Options {
		if i == m.cursor {
			b.WriteString(selectedStyle.Render(opt.Label()) + "\n")
		} else {
			b.WriteString(normalStyle.Render(opt.Label()) + "\n")
		}
	}

	return b.String()
}
