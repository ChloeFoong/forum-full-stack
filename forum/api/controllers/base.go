package controllers

import (
	"log"

	"github.com/ChloeFoong/forum/api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup of server - database and routes

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer() *Server {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Topic{},
		&models.Tag{},
	)
	models.SeedTopics(db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	server := &Server{
		DB:     db,
		Router: router,
	}

	server.InitializeRoutes()

	return server
}

func (server *Server) Run(addr string) {
	server.Router.Run(addr)
}
