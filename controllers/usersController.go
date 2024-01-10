package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"go-crud/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name     string     `json:"name" binding:"required"`
	Email    *string    `json:"email" binding:"required"`
	Birthday *time.Time `json:"birthday" binding:"required"`
}

func UsersCreate(c *gin.Context) {
	// Get Data of Request Body
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Result(c, http.StatusBadRequest, "Fail", err.Error())
	}

	// Create Post
	user := models.User{Name: input.Name, Email: input.Email, Birthday: input.Birthday, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	config.DB.Create(&user)

	utils.Result(c, http.StatusOK, "Success", user)
}

func UserDetail(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.Result(c, http.StatusNotFound, "Fail", err.Error())
	}

	utils.Result(c, http.StatusOK, "Success", user)

}

type UpdateUserInput struct {
	Name string `json:"name" binding:"required"`
}

func UserUpdate(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.Result(c, http.StatusNotFound, "Fail", "record not found")
	}

	var input UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Result(c, http.StatusBadRequest, "Fail", err.Error())
	}

	updatedUser := models.User{Name: input.Name, UpdatedAt: time.Now()}

	config.DB.Model(&user).Updates(&updatedUser)
	utils.Result(c, http.StatusOK, "Success", user)

}
