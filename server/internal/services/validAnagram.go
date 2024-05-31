package services

import "rpc-server/internal/models"

func ValidAnagram(req models.Request) models.Response {
	var res models.Response
	res.ID = req.ID

	if len(req.Params) != 2 {
		res.Error = "Invalid parameter count"
		return res
	}

	str1, str1Ok := req.Params[0].(string)
	str2, str2Ok := req.Params[1].(string)
	if str1Ok && str2Ok {
		res.Result = validAnagram(str1, str2)
		res.ResultType = "bool"
	} else {
		res.Error = "Invalid parameter type"
	}

	return res
}

func validAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[rune]int)
	for _, c := range str1 {
		count[c]++
	}
	for _, c := range str2 {
		if count[c] == 0 {
			return false
		}
		count[c]--
	}
	return true
}
