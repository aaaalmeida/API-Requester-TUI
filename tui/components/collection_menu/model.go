package collection_menu

import (
	"api-requester/context"
	"api-requester/domain/collection"
)

const WIDTH int = 30
const HEIGHT int = 25
const PADDING int = 0
const DEFAULT_COLOR string = "0" //black
const FOCUS_COLOR string = "9"   //red

// TODO: ajust based on config dotfile
type model struct {
	context        *context.AppContext
	collections    []*collection.Collection
	openCloseIndex []bool
	cursor         cursor
	width          int
	height         int
	padding        int
	default_Color  string // support ansi 16, ansi 256 and hexcode
	focus_Color    string // support ansi 16, ansi 256 and hexcode
}

type cursor struct {
	colIndex int
	reqIndex *int // if nil then points to collection, else points to requests
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:        ctx,
		collections:    nil,
		openCloseIndex: nil,
		cursor:         cursor{colIndex: 0, reqIndex: nil},
		width:          WIDTH,
		height:         HEIGHT,
		padding:        PADDING,
		default_Color:  DEFAULT_COLOR,
		focus_Color:    FOCUS_COLOR,
	}
}
