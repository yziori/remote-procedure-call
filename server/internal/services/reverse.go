package services

import (
	"rpc-server/internal/models"
)

// 文字列として入力を受け取り、逆である新しい文字列を返す
func Reverse(req models.Request) models.Response {
	var res models.Response
	res.ID = req.ID

	if len(req.Params) != 1 {
		res.Error = "Invalid parameter count"
		return res
	}

	x, ok := req.Params[0].(string)
	if ok {
		runes := []rune(x)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		res.Result = string(runes)
		res.ResultType = "string"
	} else {
		res.Error = "Invalid parameter type"
	}

	return res
}
