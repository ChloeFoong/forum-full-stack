package controllers

import (
	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-gonic/gin"
)

// handlers of the function in models for topics

func (server *Server) GetAllTopics(c *gin.Context) {
	topics, err := models.GetAllTopics(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot get all topics"})
		return
	}

	c.JSON(200, topics)
}

func (server *Server) GetTopicByID(c *gin.Context) {
	id := c.Param("id")

	var topic models.Topic

	if err := server.DB.
		Where("id = ?", id).
		First(&topic).Error; err != nil {

		c.JSON(404, gin.H{"error": "Topic not found"})
		return
	}

	c.JSON(200, topic)
}
