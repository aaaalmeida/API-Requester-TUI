package selectable

// Generic Interface. Everything that uses component MUST implements this.
type SelectOption interface {
	Label() string
	Value() any
}
