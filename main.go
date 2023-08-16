package main

import (
	"Gin_blog/dao"
	"Gin_blog/model"
)

func main() {
	user := model.User{
		Username: "LTC",
		Password: "123456",
	}
	dao.Mgr.AddUser(&user)
	dao.Mgr.FindAllUser()
}
