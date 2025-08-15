package tab

import "api-requester/context"

type model struct {
	context   *context.AppContext
	focused   bool
	activeTab int
	tabs      []TabModel
}

type TabModel struct {
	title      string
	tabContent string
	// TODO: add resto do model puxando info do request
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:   ctx,
		focused:   false,
		activeTab: 0,
		tabs: []TabModel{
			{title: "Go", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "Python", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "Java", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "JavaScript", tabContent: "qweqweqweqweqweqasdasdasd"},
		},
	}
}
