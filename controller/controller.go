package controller

import (
	"Gin_blog/dao"
	"Gin_blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(301, "/")

}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, "尝试登陆")
	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(200, "login.html", "用户不存在！")
		fmt.Println("用户不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Println("密码错误")
		} else {
			c.Redirect(301, "/")
			fmt.Println("登陆成功")
		}
	}
}

func Gologin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

/*func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}*/
