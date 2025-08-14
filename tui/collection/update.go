package collection

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m CollectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// captures msg type
	switch msg := msg.(type) {
	// User pressed a key
	case tea.KeyMsg:
		// captures which key was pressed
		switch msg.String() {

		// exit program
		case "ctrl+q", "esc":
			return m, tea.Quit

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Collections)-1 {
				m.Cursor++
			}
		case "enter":
			selectedCollection := m.Collections[m.Cursor]
			if selectedCollection.Requests == nil {
				return m, fetchRequestsFromCollectionCmd(m.Context, selectedCollection.ID)
			}
		}

		// user fetched requests from a collection
	case requestLoadedMsg:
		if msg.err != nil {
			// TODO: tratar erro
			return m, nil
		}

		/*
			FIXME: tea.Update roda em thread, talvez posso setar requests direto por indice
			com o codigo abaixo otimizando o codigo, mas caso alguma operação mude o tamanho ou
			ordem das collections isso vai resultar em bug. Testar antes de implementar mudança.
		*/
		// m.Collections[m.Cursor].Requests = msg.requests
		for i := range m.Collections {
			if m.Collections[i].ID == msg.collection_id {
				m.Collections[i].Requests = msg.requests
			}
		}
	}

	return m, nil
}
