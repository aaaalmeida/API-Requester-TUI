package main_page

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	// GENERAL NAVEGATION
	case tea.KeyMsg:
		switch msg.String() {
		// exit program
		case "ctrl+q", "esc":
			return m, tea.Quit

		// switch active component
		case "tab":
			// active component is the last
			if m.active_component_index == len(m.subcomponents)-1 {
				m.active_component_index = 0
			} else {
				m.active_component_index++
			}
		case "shift+tab":
			// active component is the first
			if m.active_component_index == 0 {
				m.active_component_index = len(m.subcomponents) - 1
			} else {
				m.active_component_index--
			}
		}

	// USER SEARCHED A COLLECTION
	case messages.LoadCollectionsMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[1], cmd = m.subcomponents[1].Update(msg)
		return m, cmd

	case messages.SendRequestToTabMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[2], cmd = m.subcomponents[2].Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.subcomponents[m.active_component_index], cmd = m.subcomponents[m.active_component_index].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
