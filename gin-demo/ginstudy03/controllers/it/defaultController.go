package it

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg":   "我是一个msg",
		"title": "我是首页",
	})
}

func (con DefaultController) News(ctx *gin.Context) {
	ctx.String(200, "新闻")
}
