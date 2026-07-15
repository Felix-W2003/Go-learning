package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func initMiddlewareOne(ctx *gin.Context) {
	fmt.Println("1-我是一个中间件--initMiddlewareOne")
	ctx.Next()

	fmt.Println("2-我是一个中间件--initMiddlewareOne")

}

func initMiddlewareTwo(ctx *gin.Context) {
	fmt.Println("1-我是一个中间件--initMiddlewareTwo")
	ctx.Next()

	fmt.Println("2-我是一个中间件--initMiddlewareTwo")

}
func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UixToTime,
	})

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	//全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)
	r.GET("/", func(ctx *gin.Context) {
		fmt.Println("这是一个首页")
		time.Sleep((time.Second * 2))
		ctx.String(http.StatusOK, "gin首页")
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页面")
	})
	r.GET("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "登录页面")
	})

	r.Run()
}
