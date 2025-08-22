package messages

import "api-requester/method"

type LoadMethodsMsg struct {
	Methods []method.Method
	Err     error
}
