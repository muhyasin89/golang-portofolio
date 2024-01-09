package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"go-crud/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostsList(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)

	utils.Result(http.StatusOK, "Success", posts)
}

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func PostsCreate(c *gin.Context) {
	// Get Data of Request Body
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Result(http.StatusBadRequest, "Fail", err.Error())
	}

	// Create Post
	post := models.Post{Title: input.Title, Content: input.Content, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	config.DB.Create(&post)

	utils.Result(http.StatusOK, "Success", post)
}

func PostDetail(c *gin.Context) {
	var post models.Post

	if err := config.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		utils.Result(http.StatusNotFound, "Fail", err.Error())
	}

	utils.Result(http.StatusOK, "Success", post)

}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func PostUpdate(c *gin.Context) {
	var post models.Post

	if err := config.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		utils.Result(http.StatusNotFound, "Fail", "record not found")
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Result(http.StatusBadRequest, "Fail", err.Error())
	}

	updatedPost := models.Post{Title: input.Title, Content: input.Content, UpdatedAt: time.Now()}

	config.DB.Model(&post).Updates(&updatedPost)
	utils.Result(http.StatusOK, "Success", post)

}

func PostDelete(c *gin.Context) {
	var post models.Post
	if err := config.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		utils.Result(http.StatusNotFound, "Fail", "record not found")
	}

	config.DB.Delete(&post)

	utils.Result(http.StatusOK, "Success", "success")
}
