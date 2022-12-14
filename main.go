package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go/data"

	_ "github.com/go-sql-driver/mysql"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"text": "hello",
	})
}

func Login(ctx *gin.Context) {
	logintext := ""
	userid, _ := ctx.GetPostForm("userId")
	password, _ := ctx.GetPostForm("password")

	userinfo := db.UserSelect(userid, password)

	if userinfo.UserId != "" {
		logintext = "ok"
	} else {
		logintext = "false"
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"text": logintext,
	})

}
