package common

type succesRes struct {
	Data   interface{} `json:"data"`
	Page   interface{} `json:"page,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccesResponse(data, page, filter interface{}) *succesRes {
	return &succesRes{Data: data, Page: page, Filter: filter}
}

func SimpleSuccesResponse(data interface{}) *succesRes {
	return NewSuccesResponse(data, nil, nil)
}
