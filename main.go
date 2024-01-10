package main

import (
	"go-crud/config"
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it works",
		})
	})

	r.GET("/posts", controllers.PostsList)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/post/:id", controllers.PostDetail)
	r.PATCH("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.POST("/users", controllers.UsersCreate)
	r.GET("/user/:id", controllers.UserDetail)
	r.PATCH("/user/:id", controllers.UserUpdate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
