package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load Template
	r.SetFuncMap(template.FuncMap{})
    r.LoadHTMLGlob("templates/*.html")

	// 핸들러
	r.GET("/", handleInit)

	r.GET("/items", handleItems)

	r.GET("/item/view/:id", handleItemViewByID)

	r.GET("/item/register", handleItemRegister)
	r.POST("/item/register", handleItemRegisterSubmit)
	r.GET("/item/register-success", handleItemRegisterSuccess)

	r.Run()
}

// handleInit 함수는 "/" --> "/items"로 리다이렉트하는 함수이다.
func handleInit(context *gin.Context) {
	context.Redirect(http.StatusSeeOther, "/items")
}