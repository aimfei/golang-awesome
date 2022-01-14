package main

import (
	"github.com/gin-gonic/gin"
	"golang-awesome/web/gin/route"
)

func main() {
	server := gin.Default()
	route.InitRoute(server)
	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
