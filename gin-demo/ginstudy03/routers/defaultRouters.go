package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/index.html", gin.H{
				"msg":   "我是一个msg",
				"title": "我是首页",
			})
		})
		defaultRouters.GET("/news", func(ctx *gin.Context) {
			ctx.String(200, "新闻")
		})
	}
}
