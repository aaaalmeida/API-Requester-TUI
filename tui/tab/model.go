package tab

import "api-requester/context"

type model struct {
	context   *context.AppContext
	activeTab int
	tabs      []TabModel
}

type TabModel struct {
	title      string
	tabContent string
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:   ctx,
		activeTab: 0,
		// TODO: add resto do model puxando info do request
		tabs: []TabModel{
			{title: "Go", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "Python", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "Java", tabContent: "qweqweqweqweqweqasdasdasd"},
			{title: "JavaScript", tabContent: "qweqweqweqweqweqasdasdasd"},
		},
	}
}
