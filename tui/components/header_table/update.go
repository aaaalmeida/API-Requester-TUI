package header_table

import (
	"api-requester/tui/messages"
	"api-requester/utils"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// INITIAL CMD
	// LOAD REQUEST
	case messages.LoadRequestMsg:
		for k, v := range msg.Request.Headers {
			keyInput := textinput.New()
			keyInput.SetValue(k)
			valueInput := textinput.New()
			valueInput.SetValue(v)
			m.rows = append(m.rows, HeaderRow{KeyTextInput: keyInput, ValueTextInput: valueInput})
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit
		case "tab":
			m.keyOrValue = !m.keyOrValue
		case "up":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		case "down":
			if m.cursorIndex < len(m.rows)-1 {
				m.cursorIndex++
			}
		case "enter":
			if m.cursorIndex == len(m.rows)-1 {
				k := textinput.New()
				k.Placeholder = "Header Key"
				v := textinput.New()
				v.Placeholder = "Header Value"
				m.rows = append(m.rows, HeaderRow{KeyTextInput: k, ValueTextInput: v})
			}
		case "delete":
			if len(m.rows) > 0 {

				m.rows = utils.RemoveFromArray(m.rows, m.cursorIndex)
			}
		}
	}

	if len(m.rows) > 0 {
		row := &m.rows[m.cursorIndex]
		if m.keyOrValue {
			row.KeyTextInput, _ = row.KeyTextInput.Update(msg)
		} else {
			row.ValueTextInput, _ = row.ValueTextInput.Update(msg)
		}
	}

	return m, nil
}
