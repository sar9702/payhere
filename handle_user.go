package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// handleSignIn 함수는 로그인 페이지를 띄우는 함수이다.
func handleSignIn(context *gin.Context) {
	context.HTML(http.StatusOK, "signin", nil)
}

// handleSignInSubmit 함수는 로그인 과정을 처리하는 함수이다.
func handleSignInSubmit(context *gin.Context) {
	id := context.PostForm("id")
	if id == "" {
		fmt.Println("ID를 입력해주세요.")
		return
	}

	password := context.PostForm("password")
	if password == "" {
		fmt.Println("비밀번호를 입력해주세요.")
		return
	}

	user, err := userByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 사용자가 입력한 비밀번호와 DB에 저장된 비밀번호가 일치하는지 확인한다.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return
	}

	// 쿠키에 Token, SignKey를 저장한다.
	context.SetCookie("SessionToken", user.Token, 3600*4, "", "", false, true)
	context.SetCookie("SessionSignKey", user.SignKey, 3600*4, "", "", false, true)

	context.Redirect(http.StatusSeeOther, "/")
}