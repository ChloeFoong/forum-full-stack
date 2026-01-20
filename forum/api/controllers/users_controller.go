package controllers

import (
	"strconv"

	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-gonic/gin"
)

// handlers of the function in models for users

func (server *Server) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(&user, server.DB); err != nil {
		c.JSON(500, gin.H{"error": "could not create account"})
		return
	}

	c.JSON(201, user)
}

func (server *Server) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required,min=3,max=30"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := models.Login(server.DB, input.Username, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
	}

	token, err := models.CreateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (server *Server) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	targetUserID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "user_id required"})
		return
	}

	if err := models.DeleteUser(server.DB, input.UserID, uint(targetUserID)); err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user deleted successfully"})
}

func (server *Server) GetUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required,username"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUser(server.DB, input.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not get user"})
		return
	}

	c.JSON(201, u)
}

func (server *Server) GetUserPosts(c *gin.Context) {
	userID := c.Param("id")

	var posts []models.Post
	if err := server.DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		c.JSON(404, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(200, posts)
}

func (server *Server) GetUserComments(c *gin.Context) {
	userID := c.Param("id")

	var comments []models.Comment
	if err := server.DB.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		c.JSON(404, gin.H{"error": "No comments found"})
		return
	}

	c.JSON(200, comments)
}
