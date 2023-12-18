package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleItems 함수는 메인 페이지를 띄우는 함수이다.
func handleItems(context *gin.Context) {
	items, err := items()
	if err != nil {
		fmt.Println(err)
	}

	switch context.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
	default:
		context.HTML(http.StatusOK, "init", gin.H{
			"products": items,
		})
	}
}

// handleItemGetByID 함수는 아이템의 상세 페이지를 띄우는 함수이다.
func handleItemGetByID(context *gin.Context) {
	id := context.Param("id")

	item, err := itemByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch context.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
	default:
		context.HTML(http.StatusOK, "item-detail", gin.H{
			"product": item,
		})
	}
}

// handleItemDeleteByID 함수는 아이템 삭제 과정을 처리하는 함수이다.
func handleItemDeleteByID(context *gin.Context) {
	id := context.Param("id")

	err := rmItem(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch context.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
	}
}

// handleItemRegister 함수는 아이템 등록 페이지를 띄우는 함수이다.
func handleItemRegister(context *gin.Context) {
	context.HTML(http.StatusOK, "register", nil)
}

// handleItemRegisterSubmit 함수는 아이템 등록 과정을 처리하는 함수이다.
func handleItemRegisterSubmit(context *gin.Context) {
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
		context.Redirect(http.StatusSeeOther, "/item/register-success")
	}
}

// handleItemRegisterSuccess 함수는 아이템 등록 완료 페이지를 띄우는 함수이다.
func handleItemRegisterSuccess(context *gin.Context) {
	context.HTML(http.StatusOK, "register-success", nil)
}