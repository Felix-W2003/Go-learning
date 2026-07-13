package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

		c.String(200, "value:%v", "hello gin")
	})

	r.GET("/old", func(c *gin.Context) {
		// Return JSON response
		c.String(200, "value:%v", "hello old")
	})

	r.GET("/index", func(c *gin.Context) {
		// Return JSON response

		c.String(200, "value:%v", "hello index")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "主要用于增加请求")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "主要用于编辑请求")
	})

	r.GET("/news", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(":8000")
}
