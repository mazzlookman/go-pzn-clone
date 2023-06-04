package formatter

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(msg string, code int, status string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Message: msg,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}
