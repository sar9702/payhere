package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// handleInit 함수는 메인 페이지를 띄우는 함수이다.
func handleInit(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "init", gin.H{
		"token": token,
	})
}

// handleItemDetail 함수는 아이템의 상세 페이지를 띄우는 함수이다.
func handleItemDetail(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "detail", gin.H{
		"token": token,
	})
}

// handleItemRegister 함수는 아이템 등록 페이지를 띄우는 함수이다.
func handleItemRegister(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "register", gin.H{
		"token": token,
	})
}

// handleItemRegisterSuccess 함수는 아이템 등록 완료 페이지를 띄우는 함수이다.
func handleItemRegisterSuccess(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "register-success", gin.H{
		"token": token,
	})
}

// handleItemEdit 함수는 아이템 수정 페이지를 띄우는 함수이다.
func handleItemEdit(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "edit", gin.H{
		"token": token,
	})
}

// handleItemEditSuccess 함수는 아이템 수정 완료 페이지를 띄우는 함수이다.
func handleItemEditSuccess(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	id := context.Param("id")
	
	context.HTML(http.StatusOK, "edit-success", gin.H{
		"token": token,
		"id": id,
	})
}

// handleSearchSubmit 함수는 URL에 검색어 데이터를 포함하여 리다이렉트하는 함수이다.
func handleSearchSubmit(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	_, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	searchWord := context.PostForm("searchWord")
	searchWord = strings.Trim(searchWord, " ")  // 앞뒤 공백 제거

	context.Redirect(http.StatusSeeOther, "/items/search?searchword=" + searchWord)
}