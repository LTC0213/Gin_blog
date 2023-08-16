package router

import (
	"Gin_blog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.Index)

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.RegisterUser)

	e.GET("/login", controller.Gologin)
	e.POST("/login", controller.Login)

	e.GET("/post_index", controller.GetPostIndex)
	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)
	e.GET("/post_detail", controller.PostDetail)
	e.Run()
}
