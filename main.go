package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load Template
	r.SetFuncMap(template.FuncMap{})
    r.LoadHTMLGlob("templates/*.html")

	// 핸들러
	r.GET("/", handleInit)

	r.GET("/register", handleRegister)

	r.Run()
}