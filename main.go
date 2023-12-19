package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	dns = "root:@tcp(localhost:3306)/payhere"
)

func main() {
	r := gin.Default()

	// Load Template and Static Files
	r.SetFuncMap(template.FuncMap{})
    r.LoadHTMLGlob("assets/templates/*.html")
	r.Static("/assets", "./assets")

	/* 핸들러 */
	r.GET("/", handleInit)

	r.GET("/items", handleItems)

	r.GET("/item/:id", handleItemGetByID)
	r.DELETE("/item/:id", handleItemDeleteByID)

	// 아이템 등록
	r.GET("/item/register", handleItemRegister)
	r.POST("/item/register", handleItemRegisterSubmit)
	r.GET("/item/register-success", handleItemRegisterSuccess)

	// 아이템 수정
	r.GET("/item/update/:id", handleItemUpdate)
	r.POST("/item/update/:id", handleItemUpdateSubmit)
	r.GET("/item/update-success/:id", handleItemUpdateSuccess)

	// 아이템 검색
	r.POST("/item/search-submit", handleItemSearchSubmit)
	r.GET("/item/search", handleItemSearch)

	// 회원가입
	r.GET("/signup", handleSignUp)
	r.POST("/signup", handleSignUpSubmit)
	r.GET("/signup-success", handleSignUpSuccess)
	r.GET("/signin", handleSignIn)
	r.POST("/signin", handleSignInSubmit)

	r.Run()
}

// handleInit 함수는 "/" --> "/items"로 리다이렉트하는 함수이다.
func handleInit(context *gin.Context) {
	context.Redirect(http.StatusSeeOther, "/items")
}