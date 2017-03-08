package db

type Meta struct {
	Total  uint `json:"total"`
	Count  uint `json:"count"`
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}

type Wrapper struct {
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta"`
}
