package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleAPIItems 함수는 DB에 저장된 모든 아이템 리스트를 반환하는 API이다.
func handleAPIItems(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	items, err := items()
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	jsonData := genResponseJson(http.StatusOK, "ok", items)
	context.JSON(http.StatusOK, jsonData)
}