package util

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
