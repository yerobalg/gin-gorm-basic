package entity

type HTTPResponse struct {
	Message   string      `json:"message"`
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data"`
}
