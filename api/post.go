package api

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts", handlers.GetPostsHandler)

	r.POST("/posts", handlers.CreatePostHandler)
}
