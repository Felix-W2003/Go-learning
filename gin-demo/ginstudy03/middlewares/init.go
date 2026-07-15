package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare(ctx *gin.Context) {
	//判断用户是否登录
	fmt.Println("这是路由分组 配置全局 中间件")
	fmt.Println(time.Now())
	fmt.Println(ctx.Request.URL)

	ctx.Set("username", "张三")
	ctxCP := ctx.Copy()
	//定义一个goroutine统计日志
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Done! in path" + ctxCP.Request.URL.Path)
	}()
}
