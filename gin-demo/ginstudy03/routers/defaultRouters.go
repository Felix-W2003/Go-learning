package routers

import (
	"gin-demo/ginstudy03/controllers/it"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", it.DefaultController{}.Index)
		defaultRouters.GET("/news", it.DefaultController{}.News)
	}
}
