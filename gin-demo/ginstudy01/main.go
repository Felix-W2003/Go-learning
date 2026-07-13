package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"score": 89,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "1111",
					Content: "2222",
				},
				&Article{
					Title:   "3333",
					Content: "4444",
				},
			},

			"news": &Article{
				Title:   "aaaa",
				Content: "bbbb",
			},
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		news := &Article{
			Title:   "this is a title ",
			Content: "this is a content",
		}
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	r.Run()
}
