package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) InitializeRoutes() {
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	//protected routes requires jwt authentication
	protected := s.Router.Group("/")
	protected.Use(AuthMiddleware())
	protected.GET("/users/:id/posts", s.GetUserPosts)
	protected.DELETE("/users/:id/delete", s.DeleteUser)
	protected.POST("/posts", s.CreatePost)
	protected.DELETE("/posts/:id/delete", s.DeletePost)
	protected.PUT("/posts/:id/update", s.UpdatePost)
	protected.POST("/posts/:id/comments", s.CreateComment)
	protected.DELETE("/comments/:id/delete", s.DeleteComment)
	protected.PUT("/comments/:id/update", s.UpdateComment)
	protected.DELETE("/tag/:id/delete", s.DeleteTag)
	protected.GET("/users/:id/comments", s.GetUserComments)

	s.Router.POST("/users", s.CreateUser)
	s.Router.POST("/login", s.Login)
	s.Router.GET("/users/:id/get", s.GetUser)

	s.Router.GET("/posts/:id/get", s.GetPost)
	s.Router.GET("/topics/:id/posts", s.GetAllPost)
	s.Router.GET("/topics/:id", s.GetTopicByID)

	s.Router.GET("/comments/:id/get", s.GetComment)
	s.Router.GET("/posts/:id/allcomment", s.GetAllComment)

	s.Router.POST("/tag", s.CreateTag)
	s.Router.GET("/tag/:id/get", s.GetTag)
	s.Router.GET("/posts/:id/alltag", s.GetAllTag)
	s.Router.GET("/topics", s.GetAllTopics)

}
