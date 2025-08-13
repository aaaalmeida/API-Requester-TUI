package collection

import (
	"api-requester/utils"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TODO: add style configuration
func (m CollectionModel) View() string {
	var response strings.Builder

	containerBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		MaxWidth(m.Width).Padding(m.Padding).
		Height(m.Height)

	selectedCollection := lipgloss.NewStyle().Bold(true).
		Background(lipgloss.Color(m.Focus_Color))

	normalCollection := lipgloss.NewStyle()

	for i, col := range m.Collections {
		string := utils.Truncate(col.Name, m.Width)

		if i == m.Cursor {
			response.WriteString(selectedCollection.Render(string))
		} else {
			response.WriteString(normalCollection.Render(string))
		}
		response.WriteRune('\n')
	}

	return containerBoxStyle.Render(response.String())
}
