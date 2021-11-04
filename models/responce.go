package models

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

const (
	ErrStatus     = "error"
	SuccessStatus = "success"
)
