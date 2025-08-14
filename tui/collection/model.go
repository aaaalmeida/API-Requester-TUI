package collection

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
type CollectionModel struct {
	Context       *context.AppContext
	Collections   []collection.Collection
	Cursor        int // Collections index
	Width         int
	Height        int
	Padding       int
	Default_Color string // support ansi 16, ansi 256 and hexcode
	Focus_Color   string // support ansi 16, ansi 256 and hexcode
}

func NewModel(ctx *context.AppContext) CollectionModel {
	collections, err := collection.GetAllCollection(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return CollectionModel{
		Context:       ctx,
		Collections:   collections,
		Cursor:        0,
		Width:         WIDTH,
		Height:        HEIGHT,
		Padding:       PADDING,
		Default_Color: DEFAULT_COLOR,
		Focus_Color:   FOCUS_COLOR,
	}
}
