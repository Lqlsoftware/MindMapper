package utils

import "encoding/json"

func GetJsonResult(msg string, code int, data interface{}) string {
	res, err := json.Marshal(map[string]interface{}{
		"msg": msg,
		"code": code,
		"data": data,
	})
	if err != nil {
		return ""
	}
	return string(res)
}