package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleRegister 함수는 아이템 등록 페이지를 띄우는 함수이다.
func handleRegister(context *gin.Context) {
	context.HTML(http.StatusOK, "register", nil)
}