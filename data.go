package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string `gorm:"unique"`
	Password string `json:"password"`
	Post     []Post
	Liked    []Liked
	Comments []Comments
}
type Liked struct {
	ID     uint
	UserID uint
	PostID uint
	User   User
	Post   Post
}
type Post struct {
	ID         uint
	Content    string
	LikeNum    int
	CommentNum int
	CreateTime string
	UserID     uint
	User       User
	Liked      []Liked
	Comments   []Comments
}
type Comments struct {
	ID         uint
	Content    string
	CreateTime string
	PostId     uint
	UserId     uint
	User       User
	Post       Post
}

var DB *gorm.DB

func init() {
	const (
		username = "root"
		password = ""
		host     = ""
		port     = 3306
		Dbname   = "forum"
		timeout  = "10s"
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败" + err.Error())

	}
	DB = db
	err = DB.AutoMigrate(&User{}, &Post{}, &Liked{}, &Comments{})
	if err != nil {
		panic(err)
	}

}
func setUser(username, password string) error {
	var user User
	user.Name = username
	user.Password = password
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func getUser(username string) (User, error) {
	var user User
	err := DB.Where("name=?", username).Take(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil

}
func savePost(contentText string, username string) {
	var post Post
	var user User
	post.Content = contentText
	post.CommentNum = 0
	post.LikeNum = 0
	post.CreateTime = currentTime()
	{ //可进行优化
		DB.Where("name=?", username).Take(&user)
		post.User.ID = user.ID
		DB.Create(&post)
	}

}
func getAllPosts(post *[]Post) {
	DB.Preload("User").Find(&post)
}
func getPost(post *Post) {
	DB.Preload("User").Preload("Comments.User").First(&post)
}
func likePost(username string, postId int, action string, likeInfo *likeInfo) {
	var liked Liked
	var post Post
	var user User
	DB.Where("name=?", username).Take(&user)
	DB.Take(&post, postId)
	liked.PostID = uint(postId)
	liked.UserID = user.ID
	switch action {
	case "query":

		err := DB.Preload("User").Preload("Post").Where("post_id=? AND user_id=?", postId, user.ID).Take(&liked).Error

		if err != nil {
			//该用户未点赞
			likeInfo.IsLike = false

		} else {
			//已经点过赞了
			likeInfo.IsLike = true

		}
	case "like":
		DB.Create(&liked)
		post.LikeNum++
		DB.Save(&post)
	case "unlike":
		DB.Where("post_id=?", postId).Where("user_id=?", user.ID).Delete(&liked)
		post.LikeNum--
		DB.Save(&post)

	}

}
func saveComment(username string, commentData CommentData) {
	var comment Comments
	var user User
	var post Post
	err := DB.Where("name=?", username).Take(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	comment.UserId = user.ID
	comment.Content = commentData.Content
	comment.CreateTime = currentTime()
	comment.PostId = uint(commentData.PostId)
	//fmt.Println(comment)
	err = DB.Create(&comment).Error
	if err != nil {
		fmt.Println(err)
	}
	post.ID = uint(commentData.PostId)
	DB.Take(&post)
	post.CommentNum++
	DB.Save(&post)

}
