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

	// 메인 페이지
	r.GET("/", handleInit)

	r.Run()
}