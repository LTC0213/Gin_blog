package dao

import (
	"Gin_blog/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User
	FindAllUser()

	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:Nhss@3120@tcp(192.168.1.213:3306)/gin_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) (user model.User) {
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) FindAllUser() {
	users := []model.User{}
	mgr.db.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post {
	posts := []model.Post{}
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) (post model.Post) {
	mgr.db.First(&post, pid)
	return post
}
