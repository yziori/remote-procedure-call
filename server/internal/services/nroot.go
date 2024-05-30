package services

import (
	"math"
	"rpc-server/internal/models"
)

// Nrootは与えられたリクエストに基づいて x の n乗根を計算し、その結果をレスポンスとして返す。
func Nroot(req models.Request) models.Response {
	var res models.Response
	res.ID = req.ID

	if len(req.Params) != 2 {
		res.Error = "Invalid parameter count"
		return res
	}

	n, nOk := req.Params[0].(float64)
	x, xOk := req.Params[1].(float64)
	if nOk && xOk {
		res.Result = math.Pow(x, 1/n)
		res.ResultType = "float64"
	} else {
		res.Error = "Invalid Parameter type"
	}

	return res
}
