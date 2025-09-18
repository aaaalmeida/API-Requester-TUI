package header_table

import "github.com/charmbracelet/bubbles/textinput"

type HeaderRow struct {
	KeyTextInput   textinput.Model
	ValueTextInput textinput.Model
}

type Model struct {
	rows        []HeaderRow
	cursorIndex int
	keyOrValue  bool // changes between key/value
	isFocused   bool
}

func NewModel() *Model {
	k := textinput.New()
	k.SetValue("qweqweqweqw")
	k.Placeholder = "Header Key"
	v := textinput.New()
	v.SetValue("asdasdas")
	v.Placeholder = "Header Value"

	return &Model{
		rows: []HeaderRow{
			HeaderRow{KeyTextInput: k, ValueTextInput: v},
		},
		isFocused:   false,
		cursorIndex: 0,
		keyOrValue:  true,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
