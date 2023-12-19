package main

// genResponseJson 함수는 API 호출 시 반환되어야 할 데이터를 반환하는 함수이다.
func genResponseJson(code int, msg string, data interface{}) map[string]interface{} {
	mapData := map[string]interface{}{
		"meta": map[string]interface{}{
			"code": code,
			"message": msg,
		},
		"data": data,
	}

	return mapData
}