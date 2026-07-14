package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func UixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UixToTime,
		"Println":    Println,
	})

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
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
			"date": 1783992427,
			"s1":   "nihao",
			"s2":   "gin",
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
