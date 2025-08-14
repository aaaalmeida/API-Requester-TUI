package collection

import (
	"api-requester/utils"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

// TODO: add style configuration
func (m CollectionModel) View() string {
	selectedCollection := lipgloss.NewStyle().Bold(true).
		Background(lipgloss.Color(m.Focus_Color))
	normalCollection := lipgloss.NewStyle()

	t := tree.New()
	for i, col := range m.Collections {
		colName := utils.Truncate(col.Name, 20)

		var subTree *tree.Tree
		if i == m.Cursor {
			subTree = t.Child(selectedCollection.Render(colName))
		} else {
			subTree = t.Child(normalCollection.Render(colName))
		}

		// only loads fetched collections
		if len(col.Requests) > 0 {
			for _, r := range col.Requests {
				subTree.Child(tree.New().Child(
					utils.Truncate(
						utils.Concatenate(strconv.Itoa(r.Method_id), r.Name), 20)))
			}
		}
	}

	containerBoxStyle := lipgloss.NewStyle().
		Height(m.Height).Width(m.Width).Padding(m.Padding).
		Border(lipgloss.ThickBorder())

	return containerBoxStyle.Render(t.String())
}
