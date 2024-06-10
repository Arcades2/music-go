package api

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.Engine) {
	r.POST("/likes", handlers.CreateLikeHandler)
}
