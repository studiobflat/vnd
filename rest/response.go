package rest

type Result struct {
	RequestId  string      `json:"requestId,omitempty"`
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total int `json:"total"`
}

type DocResult[T any] struct {
	RequestId  string      `json:"requestId,omitempty"`
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
