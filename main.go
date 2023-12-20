package main

import (
	"html/template"

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
	r.GET("/items/search", handleInit)
	r.POST("/items/search", handleSearchSubmit)
	r.GET("/item/detail", handleItemDetail)
	r.GET("/item/register", handleItemRegister)
	r.GET("/item/register-success", handleItemRegisterSuccess)
	r.GET("/item/edit", handleItemEdit)
	r.GET("/item/edit-success/:id", handleItemEditSuccess)

	r.GET("/signup", handleSignUp)
	r.POST("/signup", handleSignUpSubmit)
	r.GET("/signup-success", handleSignUpSuccess)
	r.GET("/signin", handleSignIn)
	r.POST("/signin", handleSignInSubmit)
	r.POST("/signout", handleSignOut)

	// Rest API
	r.GET("/api/items", handleAPIItems)

	r.POST("/api/item", handleAPIItemRegister)
	r.GET("/api/item/:id", handleAPIItemByID)
	r.PUT("/api/item/:id", handleAPIItemUpdateByID)
	r.DELETE("/api/item/:id", handleAPIItemDeleteByID)

	r.Run()
}