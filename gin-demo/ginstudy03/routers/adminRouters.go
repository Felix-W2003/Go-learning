package routers

import "github.com/gin-gonic/gin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(ctx *gin.Context) {
			ctx.String(200, "用户列表")
		})
		adminRouters.GET("/user/add", func(ctx *gin.Context) {
			ctx.String(200, "添加用户")
		})
		adminRouters.GET("/user/edit", func(ctx *gin.Context) {
			ctx.String(200, "修改")
		})
		adminRouters.GET("/article", func(ctx *gin.Context) {
			ctx.String(200, "新闻列表")
		})
	}
}
