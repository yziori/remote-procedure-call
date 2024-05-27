package services

import (
	"math"
	"rpc-server/internal/models"
)

// Floor returns the greatest integer value less than or equal to x.
func Floor(req models.Request) models.Response {
	var res models.Response
	res.ID = req.ID

	if len(req.Params) == 1 {
		x, ok := req.Params[0].(float64)
		if ok {
			res.Result = math.Floor(x)
			res.ResultType = "int"
		} else {
			res.Error = "Invalid params"
		}
	} else {
		res.Error = "Invalid params"
	}

	return res
}
