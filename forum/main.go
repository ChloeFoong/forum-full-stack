package main

import (
	"github.com/ChloeFoong/forum/api/controllers"
)

func main() {
	server := controllers.NewServer()
	server.Run(":8080")

}
