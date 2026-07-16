package admin

import (
	"fmt"
	"gin-demo/ginstudy03/models"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//	c.String(200, "用户列表")
	//con.success(c)

	//查询数据库
	// userList := []models.User{}
	// models.DB.Find(&userList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": userList,
	// })

	//查询age>20的用户
	userList := []models.User{}
	models.DB.Where("age>20").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

func (con UserController) Add(c *gin.Context) {
	// c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
	user := models.User{
		Id:       4,
		Username: "许立军",
		Age:      22,
		Email:    "12123@qq.com",
		AddTime:  1000002,
	}

	models.DB.Create(&user)
	fmt.Println(user)
	c.String(200, "success add")
}

func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	dst := path.Join("./static/upload", file.Filename)
	//file.Filename
	if err == nil {

		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"sucess":   true,
			"username": username,
			"dst":      dst,
		})
	} else {
		c.JSON(400, gin.H{
			"success":  false,
			"username": username,
			"dst":      dst,
		})
	}
	log.Println(file.Filename)
	c.String(200, "执行上传")

}

func (con UserController) MutiAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/usermutiadd.html", gin.H{})
}
func (con UserController) DoMutiUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["face[]"]
	username := c.PostForm("username")
	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"sucess":   true,
		"username": username,
	})
}
func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")
	face1, err1 := c.FormFile("face1")
	dst1 := path.Join("./static/upload", face1.Filename)
	//file.Filename
	if err1 == nil {

		c.SaveUploadedFile(face1, dst1)
		c.JSON(http.StatusOK, gin.H{
			"sucess":   true,
			"username": username,
			"dst":      dst1,
		})
	} else {
		c.JSON(400, gin.H{
			"success":  false,
			"username": username,
			"dst":      dst1,
		})
	}
	face2, err2 := c.FormFile("face2")
	dst2 := path.Join("./static/upload", face2.Filename)
	//file.Filename
	if err2 == nil {

		c.SaveUploadedFile(face2, dst2)
		c.JSON(http.StatusOK, gin.H{
			"sucess":   true,
			"username": username,
			"dst":      dst2,
		})
	} else {
		c.JSON(400, gin.H{
			"success":  false,
			"username": username,
			"dst":      dst2,
		})
	}
	c.String(200, "执行上传")
}
