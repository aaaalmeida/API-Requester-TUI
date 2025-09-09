package utils

func TransformMethodIdToVerb(id int) string {
	var verb string
	switch id {
	case 1:
		verb = "GET"
	case 2:
		verb = "POST"
	case 3:
		verb = "PUT"
	case 4:
		verb = "DELETE"
	case 5:
		verb = "PATCH"
	case 6:
		verb = "HEAD"
	case 7:
		verb = "TRACE"
	case 8:
		verb = "OPTIONS"
	}
	return verb
}
