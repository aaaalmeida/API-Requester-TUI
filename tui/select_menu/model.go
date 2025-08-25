package select_menu

type Model struct {
	options      []selectOption
	cursor       int
	selectedItem int
	isOpened     bool
	width        int
}

// Generic Interface. Everything that uses component MUST implements this.
type selectOption interface {
	Label() string
	Value() any
}

func NewModel() Model {
	return Model{
		cursor:       0,
		selectedItem: 0,
		options:      nil,
		isOpened:     false,
		width:        20,
	}
}
