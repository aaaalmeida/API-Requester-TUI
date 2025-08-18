package collection_menu

import (
	cmd "api-requester/tui/collection_menu/command"
	collection_menu "api-requester/tui/collection_menu/msg"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// captures msg type
	switch msg := msg.(type) {
	// User pressed a key
	case tea.KeyMsg:
		// captures which key was pressed
		switch msg.String() {

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.collections)-1 {
				m.cursor++
			}
		case "enter":
			// alternate between open/close collection in menu
			m.openCloseIndex[m.cursor] = !m.openCloseIndex[m.cursor]

			selectedCollection := m.collections[m.cursor]
			if selectedCollection.Requests == nil {
				return m, cmd.FetchRequestsFromCollectionCmd(m.context, selectedCollection.ID)
			}
		}

		// user fetched requests from a collection
	case collection_menu.LoadCollectionMsg:
		if msg.Err != nil {
			// TODO: tratar erro
			return m, nil
		}

		/*
			FIXME: tea.Update roda em thread, talvez posso setar requests direto por indice
			com o codigo abaixo otimizando o codigo, mas caso alguma operação mude o tamanho ou
			ordem das collections isso vai resultar em bug. Testar antes de implementar mudança.
		*/
		// m.Collections[m.Cursor].Requests = msg.requests
		for i := range m.collections {
			if m.collections[i].ID == msg.Collection_id {
				m.collections[i].Requests = msg.Requests
			}
		}
	}

	return m, nil
}
