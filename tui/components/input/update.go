package input

import (
	cmd "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, cmd.SendStringFromInputCmd(m.textInput.Value())
		}
	}

	var c tea.Cmd
	m.textInput, c = m.textInput.Update(msg)
	cmds = append(cmds, c)
	return m, tea.Batch(cmds...)
}
