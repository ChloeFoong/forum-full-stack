package controllers

import (
	"strconv"

	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-gonic/gin"
)

// handlers of the function in models for tags

func (server *Server) CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateTag(server.DB, &tag); err != nil {
		c.JSON(500, gin.H{"error": "could not create account"})
		return
	}

	c.JSON(201, tag)
}

func (server *Server) DeleteTag(c *gin.Context) {
	idParam := c.Param("id")
	tagID, _ := strconv.ParseUint(idParam, 10, 32)
	var input struct {
		UserID uint `json:"user_id"`
	}
	c.ShouldBindJSON(&input)

	tag, err := models.GetTag(server.DB, uint(tagID))
	if err != nil {
		c.JSON(404, gin.H{"error": "tag not found"})
		return
	}

	if err := models.DeleteTag(server.DB, input.UserID, tag); err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "tag deleted successfully"})
}

func (server *Server) GetTag(c *gin.Context) {
	var input struct {
		TagID uint `json:"tag_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	t, err := models.GetPost(server.DB, input.TagID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not get tag"})
		return
	}

	c.JSON(201, t)
}

func (server *Server) GetAllTag(c *gin.Context) {
	postParam := c.Param("id")
	postID, err := strconv.ParseUint(postParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid post id"})
		return
	}

	tags, err := models.GetAllTag(server.DB, uint(postID))
	if err != nil {
		c.JSON(404, gin.H{"error": "no tags found for this post"})
		return
	}

	c.JSON(200, tags)
}
