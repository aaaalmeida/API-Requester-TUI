package collection_menu

import (
	cmd "api-requester/tui/commands"
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// captures msg type
	switch msg := msg.(type) {
	// USER PRESSED A KEY
	case tea.KeyMsg:
		// captures which key was pressed
		switch msg.String() {
		case "up":
			items := m.visibleItems()
			for i := range items {
				if items[i].Equal(m.cursor) {
					if i > 0 {
						m.cursor = items[i-1]
					}
					break
				}
			}

		case "down":
			items := m.visibleItems()
			for i := range items {
				if items[i].Equal(m.cursor) {
					if i < len(items)-1 {
						m.cursor = items[i+1]
					}
					break
				}
			}

		case "enter", " ":
			if m.cursor.reqIndex == nil {
				idx := m.cursor.colIndex
				m.openCloseIndex[idx] = !m.openCloseIndex[idx]

				selectedCollection := m.collections[idx]
				if selectedCollection.Requests == nil {
					return m, cmd.FetchRequestsFromCollectionCmd(m.context, selectedCollection.ID)
				}
			} else {
				// TODO: send request to tab

			}
		}

	// INITIAL UPDATE
	case messages.LoadCollectionsMsg:
		m.collections = msg.Collections
		m.openCloseIndex = make([]bool, len(m.collections))
		return m, nil

	// USER FETCHED REQUESTS FROM A COLLECTION
	case messages.LoadRequestFromCollectionMsg:
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

// Controls tree navigation
func (m model) visibleItems() []cursor {
	var items []cursor
	for i, col := range m.collections {
		items = append(items, cursor{colIndex: i, reqIndex: nil}) // collection position
		if m.openCloseIndex[i] && col.Requests != nil {
			for j := range col.Requests {
				jCopy := j
				items = append(items, cursor{colIndex: i, reqIndex: &jCopy}) // request position
			}
		}
	}
	return items
}

// Compares value, not address
func (c cursor) Equal(other cursor) bool {
	if c.colIndex != other.colIndex {
		return false
	}
	if c.reqIndex == nil && other.reqIndex == nil {
		return true
	}
	if c.reqIndex != nil && other.reqIndex != nil {
		return *c.reqIndex == *other.reqIndex
	}
	return false
}
