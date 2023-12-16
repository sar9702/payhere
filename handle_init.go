package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// handleInit 함수는 메인 페이지를 띄우는 함수이다.
func handleInit(context *gin.Context) {
	context.HTML(http.StatusOK, "init", gin.H{
		"messages": "1234",
	})
}