package header_table

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m Model) View() string {
	var headers [][]string
	for _, r := range m.rows {
		headers = append(headers, []string{r.KeyTextInput.Value(), r.ValueTextInput.Value()})
	}

	return table.New().Border(lipgloss.NormalBorder()).
		Headers("Key", "Value").
		Rows(headers...).Render()
}
