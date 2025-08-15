package collection_menu

import (
	"api-requester/utils"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

// TODO: add style configuration
func (m model) View() string {
	selectedCollection := lipgloss.NewStyle().Bold(true).
		Background(lipgloss.Color(m.focus_Color))
	normalCollection := lipgloss.NewStyle()

	t := tree.New()
	for i, col := range m.collections {
		colName := utils.Truncate(col.Name, 20)

		var subTree *tree.Tree
		if i == m.cursor {
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
		Height(m.height).Width(m.width).Padding(m.padding).
		Border(lipgloss.ThickBorder())

	return containerBoxStyle.Render(t.String())
}
