package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-gonic/gin"
)

// handlers of the function in models for posts

func (server *Server) CreatePost(c *gin.Context) {
	var input struct {
		Heading string       `json:"heading"`
		Content string       `json:"content"`
		Tag     []models.Tag `json:"tag"`
		TopicID uint         `json:"topic_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usernameVal, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	username := usernameVal.(string)

	var user models.User
	if err := server.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	post := models.Post{
		Heading: input.Heading,
		Content: input.Content,
		Tag:     input.Tag,
		UserID:  user.ID,
		TopicID: input.TopicID,
	}

	if err := server.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (server *Server) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	postID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid post id"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	var post models.Post
	if err := server.DB.First(&post, postID).Error; err != nil {
		c.JSON(404, gin.H{"error": "post not found"})
		return
	}

	userID := post.UserID

	updatedPost, err := models.UpdatePost(server.DB, userID, input.Title, input.Content, post.ID)

	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedPost)
}

func (server *Server) DeletePost(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid post ID"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Missing Authorization header"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	username, err := models.VerifyTokenAndGetUsername(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	user, err := models.GetUser(server.DB, username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	post, err := models.GetPost(server.DB, uint(postID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != user.ID {
		c.JSON(403, gin.H{"error": "Cannot delete another user's post"})
		return
	}

	if err := models.DeletePost(server.DB, user.ID, post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

func (server *Server) GetPost(c *gin.Context) {

	var input struct {
		PostID uint `json:"post_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	p, err := models.GetPost(server.DB, input.PostID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not get user"})
		return
	}

	c.JSON(201, p)
}

func (server *Server) GetAllPost(c *gin.Context) {
	topicParam := c.Param("id")
	topicID, err := strconv.ParseUint(topicParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid topic id"})
		return
	}

	posts, err := models.GetAllPost(server.DB, uint(topicID))
	if err != nil {
		c.JSON(404, gin.H{"error": "no posts found for this topic"})
		return
	}

	c.JSON(200, posts)
}
