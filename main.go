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

	r.GET("/posts", controllers.PostsList)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts/:id", controllers.PostDetail)
	r.PATCH("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
