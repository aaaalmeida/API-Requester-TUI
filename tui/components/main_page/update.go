package main_page

import (
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.activeComponentIndex == len(m.subcomponents)-1 {
				m.activeComponentIndex = 0
			} else {
				m.activeComponentIndex++
			}
		case "shift+tab":
			// active component is the first
			if m.activeComponentIndex == 0 {
				m.activeComponentIndex = len(m.subcomponents) - 1
			} else {
				m.activeComponentIndex--
			}
		}

	// USER SEARCHED A COLLECTION
	case messages.LoadCollectionsMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[CollectionMenuIndex], cmd = m.subcomponents[CollectionMenuIndex].Update(msg)
		return m, cmd

	// SEND A REQUEST FROM COLLECTION_MENU TO HEADER
	case messages.SendRequestMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[HeaderIndex], cmd = m.subcomponents[HeaderIndex].Update(msg)
		return m, cmd

		// FETCH METHODS FROM DB TO REQUEST_HEADER_BODY
	case messages.LoadMethodsMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}

		var cmd tea.Cmd
		m.subcomponents[RequestHeadersIndex], cmd = m.subcomponents[RequestHeadersIndex].Update(msg)
		return m, cmd

		// SEND A REQUEST FROM HEADER TO MAIN MENU
	case messages.LoadRequestMsg:
		if msg.Err != nil {
			// FIXME: TRATAR ERRO
			return m, nil
		}
		var cmd tea.Cmd
		m.subcomponents[RequestHeadersIndex], cmd = m.subcomponents[RequestHeadersIndex].Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.subcomponents[m.activeComponentIndex], cmd = m.subcomponents[m.activeComponentIndex].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
