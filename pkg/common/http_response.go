package common

type successRes struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Paging  interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, paging interface{}) *successRes {
	return &successRes{Data: data, Paging: paging}
}

func SimpleDataSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data}
}

func SimpleMessageSuccessResponse(message interface{}) *successRes {
	return &successRes{Message: message}
}
