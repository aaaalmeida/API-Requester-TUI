package utils

func TransformMethodIdToVerbColored(id int) string {
	var Reset = "\033[0m"

	var Black = "\033[30m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"

	var verb string
	switch id {
	default:
		verb = Green + "GET"
	case 2:
		verb = Blue + "POST"
	case 3:
		verb = Yellow + "PUT"
	case 4:
		verb = Red + "DELETE"
	case 5:
		verb = Magenta + "PATCH"
	case 6:
		verb = Cyan + "HEAD"
	case 7:
		verb = Gray + "TRACE"
	case 8:
		verb = Black + "OPTIONS"
	}

	verb += Reset
	return verb
}
