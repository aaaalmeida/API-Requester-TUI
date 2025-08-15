package collection_menu

import (
	"api-requester/collection"
	"api-requester/context"
	"log"
)

const WIDTH int = 30
const HEIGHT int = 25
const PADDING int = 0
const DEFAULT_COLOR string = "0" //black
const FOCUS_COLOR string = "9"   //red

// TODO: ajust based on config dotfile
type model struct {
	context       *context.AppContext
	collections   []collection.Collection
	cursor        int // Collections index
	width         int
	height        int
	padding       int
	default_Color string // support ansi 16, ansi 256 and hexcode
	focus_Color   string // support ansi 16, ansi 256 and hexcode
}

func NewModel(ctx *context.AppContext) model {
	collections, err := collection.GetAllCollection(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return model{
		context:       ctx,
		collections:   collections,
		cursor:        0,
		width:         WIDTH,
		height:        HEIGHT,
		padding:       PADDING,
		default_Color: DEFAULT_COLOR,
		focus_Color:   FOCUS_COLOR,
	}
}
