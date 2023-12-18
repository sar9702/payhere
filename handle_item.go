package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleRegister 함수는 아이템 등록 페이지를 띄우는 함수이다.
func handleRegister(context *gin.Context) {
	context.HTML(http.StatusOK, "register", nil)
}

// handleRegisterSubmit 함수는 아이템 등록 과정을 처리하는 함수이다.
func handleRegisterSubmit(context *gin.Context) {
	var item Item

	// POST request 데이터를 item에 넣는다.
	err := context.Bind(&item)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = addItem(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch context.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
	default:
		context.Redirect(http.StatusSeeOther, "/register-success")
	}
}

// handleRegisterSuccess 함수는 아이템 등록 완료 페이지를 띄우는 함수이다.
func handleRegisterSuccess(context *gin.Context) {
	context.HTML(http.StatusOK, "register-success", nil)
}