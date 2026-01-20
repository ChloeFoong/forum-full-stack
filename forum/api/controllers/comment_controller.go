package controllers

import (
	"net/http"
	"strconv"

	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-gonic/gin"
)

// handlers of the function in models for comments

func (server *Server) CreateComment(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := models.VerifyTokenAndGetUsername(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := models.GetUser(server.DB, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		PostID  uint   `json:"post_id"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := models.CreateComment(server.DB, user.ID, input.PostID, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (server *Server) UpdateComment(c *gin.Context) {
	idParam := c.Param("id")
	commentID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid comment id"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "missing token"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := models.VerifyTokenAndGetUsername(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}

	user, err := models.GetUser(server.DB, username)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	var input struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	updatedComment, err := models.UpdateComment(
		server.DB,
		user.ID,
		uint(commentID),
		input.Content,
	)

	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedComment)
}

func (server *Server) DeleteComment(c *gin.Context) {
	idParam := c.Param("id")
	commentID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid comment id"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "missing token"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := models.VerifyTokenAndGetUsername(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}

	user, err := models.GetUser(server.DB, username)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	comment, err := models.GetComment(server.DB, uint(commentID))
	if err != nil {
		c.JSON(404, gin.H{"error": "comment not found"})
		return
	}

	if err := models.DeleteComment(server.DB, user.ID, comment); err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "comment deleted successfully"})
}

func (server *Server) GetComment(c *gin.Context) {
	var input struct {
		CommentID uint `json:"comment_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment, err := models.GetComment(server.DB, input.CommentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not get user"})
		return
	}

	c.JSON(201, comment)
}

func (server *Server) GetAllComment(c *gin.Context) {
	postParam := c.Param("id")
	postID, err := strconv.ParseUint(postParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid topic id"})
		return
	}

	comments, err := models.GetAllComment(server.DB, uint(postID))
	if err != nil {
		c.JSON(404, gin.H{"error": "no comments found for this post"})
		return
	}

	c.JSON(200, comments)
}
