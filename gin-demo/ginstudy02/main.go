package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
		})
		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻",
		})
	})

	r.GET("/article", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情页面",
			"id":  id,
		})
	})

	r.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	r.POST("/doAddUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "20")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/getuser", func(ctx *gin.Context) {
		user := &UserInfo{}

		if err := ctx.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	r.POST("/doAddUser2", func(ctx *gin.Context) {
		user := &UserInfo{}

		if err := ctx.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.DELETE("/delete", func(ctx *gin.Context) {
		//article := &Article{}
		xmlSliceData, _ := ctx.GetRawData()
		fmt.Println(xmlSliceData)
		article := &Article{}
		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			fmt.Printf("%#v", article)
			ctx.JSON(http.StatusOK, article)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	r.GET("/list/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(200, "%v", cid)
	})

	r.Run()
}
