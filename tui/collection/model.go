package collection

import (
	"api-requester/collection"
	"api-requester/context"
	"log"
)

// TODO: ajust size based on config dotfile
const WIDTH int = 30
const HEIGHT int = 30
const PADDING int = 0
const DEFAULT_COLOR string = "0" //black
const FOCUS_COLOR string = "9"   //red

type CollectionModel struct {
	Collections   []collection.Collection
	Cursor        int
	Width         int
	Height        int
	Padding       int
	Default_Color string // support ansi 16, ansi 256 and hexcode
	Focus_Color   string // support ansi 16, ansi 256 and hexcode
}

// func NewModel(collections []collection.Collection) CollectionModel {
func NewModel(ctx *context.AppContext) CollectionModel {
	collections, err := collection.GetAllCollection(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return CollectionModel{
		Collections:   collections,
		Cursor:        0,
		Width:         WIDTH,
		Height:        HEIGHT,
		Padding:       PADDING,
		Default_Color: DEFAULT_COLOR,
		Focus_Color:   FOCUS_COLOR,
	}
}
