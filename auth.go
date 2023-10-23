package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Info struct {
	Username string `form:"Username" binding:"required"`
	Password string `form:"Password" binding:"required"`
}
type FormData struct {
	ContentText string `form:"contentText" binding:"required"`
}

type CommentData struct {
	PostId  int    `json:"postId"`
	Content string `json:"content"`
}
type likeInfo struct {
	IsLike  bool `json:"isLike"`
	IsGuest bool `json:"isGuest"`
}

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("isLogin")
		if err != nil || cookie == "no" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()

		}
	}
}
func _error(c *gin.Context) {
	c.JSON(404, gin.H{"err": "404"})
}
func form(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}
func formData(c *gin.Context) {
	var formData FormData
	err := c.Bind(&formData)
	if err != nil {
		fmt.Println(err)
	}
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(301, "/error")
		return
	}
	savePost(formData.ContentText, username)
	c.Redirect(301, "/")

}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func getData(c *gin.Context) {
	var posts []Post
	getAllPosts(&posts)
	data, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}
	c.Data(http.StatusOK, "application/json;charset=utf-8", data)
}

func loginData(c *gin.Context) {
	var info Info
	if err := c.Bind(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	err := logIn(info)
	if err != nil {
		if err.Error() == "密码错误" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请先注册",
			})
		}
	}

	c.SetCookie("username", info.Username, 3600, "/", "localhost", false, false)
	c.SetCookie("isLogin", "yes", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功"})

}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}
func logout(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", "localhost", false, false)
	c.SetCookie("isLogin", "", -1, "/", "localhost", false, true)
}
func reg(c *gin.Context) {
	c.HTML(http.StatusOK, "reg.html", nil)

}
func regData(c *gin.Context) {

	var info Info
	if err := c.Bind(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := register(info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
func postLike(c *gin.Context) {
	var action struct {
		Action string `json:"action"`
		PostId string `json:"postId"`
	}
	var likeInfo likeInfo
	if err := c.ShouldBindJSON(&action); err != nil {
		c.Redirect(301, "/error")
		return
	}

	postId, err := strconv.Atoi(action.PostId)

	if err != nil {
		fmt.Println(err)
		return
	}
	username, err := c.Cookie("username")
	if err != nil {
		likeInfo.IsGuest = true
		c.JSON(http.StatusOK, likeInfo)
		return
	}

	likePost(username, postId, action.Action, &likeInfo)
	c.JSON(http.StatusOK, likeInfo)
}
func comment(c *gin.Context) {
	c.HTML(http.StatusOK, "comment.html", nil)
}
func loadComment(c *gin.Context) {
	var post Post
	postId := c.Param("postId")
	postID, err := strconv.Atoi(postId)
	post.ID = uint(postID)
	if err != nil {
		fmt.Println(err)
	}
	getPost(&post)
	c.JSON(http.StatusOK, post)
}
func updateComment(c *gin.Context) {
	var commentData CommentData
	err := c.ShouldBindJSON(&commentData)
	//fmt.Println("这是评论数据")
	//fmt.Println(commentData)
	if err != nil {
		fmt.Println(err)
	}
	username, err := c.Cookie("username")
	if err != nil {
		fmt.Println(err)
	}
	saveComment(username, commentData)
}
