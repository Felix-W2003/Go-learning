package main

import (
	"gin-demo/ginstudy03/models"
	"gin-demo/ginstudy03/routers"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UixToTime,
	})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	r.Use()
	//配置session中间件
	//创建基于cookie的存储引擎，secret123参数用于加密的密钥
	//store := cookie.NewStore([]byte("secret123"))

	//创建基于redis的存储引擎
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	// defaultRouters := r.Group("/")
	// {
	// 	defaultRouters.GET("/", func(ctx *gin.Context) {
	// 		ctx.String(200, "首页")
	// 	})
	// 	defaultRouters.GET("/news", func(ctx *gin.Context) {
	// 		ctx.String(200, "新闻")
	// 	})
	// }

	// apiRouters := r.Group("/api")
	// {
	// 	apiRouters.GET("/", func(ctx *gin.Context) {
	// 		ctx.String(200, "我是一个api接口")
	// 	})
	// 	apiRouters.GET("/userlist", func(ctx *gin.Context) {
	// 		ctx.String(200, "我是一个api接口--userlist")
	// 	})
	// 	apiRouters.GET("/plist", func(ctx *gin.Context) {
	// 		ctx.String(200, "我是一个api接口--plist")
	// 	})
	// }
	// adminRouters := r.Group("/admin")
	// {
	// 	adminRouters.GET("/", func(ctx *gin.Context) {
	// 		ctx.String(200, "后台首页")
	// 	})
	// 	adminRouters.GET("/user", func(ctx *gin.Context) {
	// 		ctx.String(200, "用户列表")
	// 	})
	// 	adminRouters.GET("/user/add", func(ctx *gin.Context) {
	// 		ctx.String(200, "添加用户")
	// 	})
	// 	adminRouters.GET("/user/edit", func(ctx *gin.Context) {
	// 		ctx.String(200, "修改")
	// 	})
	// 	adminRouters.GET("/article", func(ctx *gin.Context) {
	// 		ctx.String(200, "新闻列表")
	// 	})
	// }
	r.Run(":80")
}
