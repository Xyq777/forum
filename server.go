package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.POST("/reg", regData)
	r.GET("/form", checkLogin(), form)
	r.POST("/form", checkLogin(), formData)
	r.GET("/", index)
	r.GET("/data", getData)
	r.GET("/reg", reg)
	r.GET("/login", login)
	r.POST("/login", loginData)
	r.GET("/error", _error)
	r.GET("/logout", checkLogin(), logout)
	r.POST("/postLike", checkLogin(), postLike)
	r.GET("/comment/:postId", checkLogin(), comment)
	r.POST("/comment/:postId", checkLogin(), loadComment)
	r.POST("/comment", checkLogin(), updateComment)
	{
		r.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"1": 1,
			})
			c.JSON(200, gin.H{
				"2": 1,
			})
		})
	}
	//r.GET("/logReceive", logReceive)
	r.Run(":8080")
}
