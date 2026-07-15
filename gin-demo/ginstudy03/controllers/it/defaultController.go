package it

import (
	"fmt"
	"gin-demo/ginstudy03/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(ctx *gin.Context) {
	fmt.Println(models.UixToTime(1784083129))

	ctx.SetCookie("username", "张三", 3600, "/", "a.wf.com", false, false)
	ctx.SetCookie("hobby", "吃饭、睡觉", 5, "/", "localhost", false, false)
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg":   "我是一个msg",
		"title": "我是首页",
		"t":     1784083129,
	})
}

func (con DefaultController) News(ctx *gin.Context) {
	username, _ := ctx.Cookie("username")
	hobby, _ := ctx.Cookie("hobby")
	ctx.String(200, "username=%v-----hobby=%v", username, hobby)
}

func (con DefaultController) Shop(ctx *gin.Context) {
	username, _ := ctx.Cookie("username")
	ctx.String(200, "cookie="+username)
}
func (con DefaultController) DeleteCookie(ctx *gin.Context) {
	ctx.SetCookie("username", "张三", -1, "/", "localhost", false, false)
	ctx.String(200, "删除成功")
}
