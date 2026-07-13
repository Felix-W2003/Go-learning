package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 定义XML数据结构体，xml标签控制节点名称
type ResMsg struct {
	Success bool   `xml:"success"`
	Msg     string `xml:"msg"`
}

func main() {
	r := gin.Default()
	//配置模版文件
	r.LoadHTMLGlob("../templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "value:%v", "index")
	})

	r.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "hello gin1",
		})
	})
	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "hello gin2",
		})
	})

	r.GET("/json3", func(ctx *gin.Context) {

		a := Article{
			Title:   "Im a title",
			Content: "text",
			Desc:    "this is describle",
		}
		ctx.JSON(200, a)
	})

	r.GET("/jsonp", func(ctx *gin.Context) {

		a := Article{
			Title:   "Im a title-jsonp",
			Content: "text",
			Desc:    "this is describle",
		}
		ctx.JSONP(200, a)
	})

	r.GET("/xml", func(ctx *gin.Context) {
		data := ResMsg{
			Success: true,
			Msg:     "this is XML",
		}
		ctx.XML(http.StatusOK, data)
	})

	r.GET("/news", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})

	r.GET("/goods", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品的数据",
		})
	})

	r.Run()
}
