package request_response_box

type Model struct {
	// request *request.Request
	body *string
}

func NewModel() Model {
	return Model{
		body: nil,
	}
}
