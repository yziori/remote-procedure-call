package models

type Request struct {
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ParamTypes []string      `json:"pram_types"`
	ID         int           `json:"id"`
}
