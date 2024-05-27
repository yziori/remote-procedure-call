package models

type Response struct {
	Result     interface{} `json:"result"`
	ResultType string      `json:"result_type"`
	ID         int         `json:"id"`
	Error      string      `json:"error,omitempty"`
}
