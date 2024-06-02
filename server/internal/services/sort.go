package services

import (
	"rpc-server/internal/models"
	"sort"
)

func Sort(req models.Request) models.Response {
	var res models.Response
	res.ID = req.ID

	if len(req.Params) != 1 {
		res.Error = "Invalid parameter count"
		return res
	}

	strArr, ok := req.Params[0].([]interface{})
	if ok {
		var strs []string
		for _, v := range strArr {
			if str, strOk := v.(string); strOk {
				strs = append(strs, str)
			} else {
				res.Error = "Invalid parameter type"
				break
			}
		}
		if res.Error == "" {
			res.Result = sortStrings(strs)
			res.ResultType = "[]string"
		}
	} else {
		res.Error = "Invalid parameter type"
	}

	return res
}

func sortStrings(strs []string) []interface{} {
	sort.Strings(strs)
	result := make([]interface{}, len(strs))
	for i, v := range strs {
		result[i] = v
	}
	return result
}
