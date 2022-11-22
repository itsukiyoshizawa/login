package main

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/", Index)
	r.POST("/login", Login)
	r.Run(":8080")
}
