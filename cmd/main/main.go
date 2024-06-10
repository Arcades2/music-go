package main

import (
	"main/api"
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api.PostRoutes(r)
	api.LikeRoutes(r)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
