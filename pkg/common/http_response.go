package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, paging interface{}) *successRes {
	return &successRes{Data: data, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil}
}
