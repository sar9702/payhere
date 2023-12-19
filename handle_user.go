package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleSignUp 함수는 회원가입 페이지를 띄우는 함수이다.
func handleSignUp(context *gin.Context) {
	context.HTML(http.StatusOK, "signup", nil)
}

// handleSignUpSubmit 함수는 회원가입 과정을 처리하는 함수이다.
func handleSignUpSubmit(context *gin.Context) {
	id := context.PostForm("id")
	password := context.PostForm("password")

	encryptedPW, err := encrypt(password)
	if err != nil{
		fmt.Println(err)
		return
	}

	user := User{
		ID: id,
		Password: encryptedPW,
	}

	err = user.CreateToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = addUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	context.Redirect(http.StatusSeeOther, "/signup-success")
}

// handleSignUpSuccess 함수는 회원가입 완료 페이지를 띄우는 함수이다.
func handleSignUpSuccess(context *gin.Context) {
	context.HTML(http.StatusOK, "signup-success", nil)
}